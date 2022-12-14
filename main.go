package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
	"karmem.org/golang"

	waaskm "waas/km"

	"waas/capabilities/wait"
)

var engine wapc.Engine
var moduleMapMutex = &sync.RWMutex{}
var moduleCtx context.Context = context.Background()
var moduleMap map[string]wapc.Module

var kmWriterPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

var managedLocations []string

func resetModules() {
	// Reset capabilities first
	wait.Reset()

	func() {
		moduleMapMutex.RLock()
		defer moduleMapMutex.RUnlock()
		for _, module := range moduleMap {
			module.Close(moduleCtx)
		}
	}()

	moduleMapMutex.Lock()
	defer moduleMapMutex.Unlock()
	moduleMap = make(map[string]wapc.Module)
}

func loadModule(name string) {
	folderName := strings.Split(name, "-")[0]
	wasm, err := os.ReadFile(fmt.Sprintf("services/%s/%s.wasm", folderName, name))
	check(err)
	module, err := engine.New(moduleCtx, host, wasm, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	check(err)

	moduleMapMutex.Lock()
	defer moduleMapMutex.Unlock()
	moduleMap[name] = module
}

func getModule(name string) (wapc.Module, error) {
	moduleMapMutex.RLock()
	defer moduleMapMutex.RUnlock()
	if module, exist := moduleMap[name]; exist {
		return module, nil
	}
	return nil, fmt.Errorf("module not loaded or not found: %s", name)
}

func main() {
	locationsFlag := flag.String("locations", "", "Comma separated of locations handled by this server")
	portFlag := flag.Int("port", 0, "Server port")
	flag.Parse()
	managedLocations = strings.Split(*locationsFlag, ",")
	port := *portFlag
	if len(managedLocations) == 0 || port == 0 {
		panic("require locations and port")
	}
	fmt.Printf("Starting WAAS server managedLocations=%v port=%d\n", managedLocations, *portFlag)

	engine = wazero.Engine()
	resetModules()

	fmt.Println("loading modules")
	for _, managedLocation := range managedLocations {
		loadModule(fmt.Sprintf("hello-%s", managedLocation))
	}
	loadModule("capitalize")
	loadModule("ping")
	fmt.Println("modules loaded")

	handler := http.HandlerFunc(handleHTTP)
	fmt.Println("Listening...")
	http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), handler)

	resetModules() // close loaded modules
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	ctx := req.Context()
	reqBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %v", err)
	}
	resBytes, err := invoke(ctx, reqBytes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %v", err)
	}
	_, err = w.Write(resBytes)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "error: %v", err)
	}
	duration := time.Since(start)
	fmt.Printf("handleHTTP took %v\n", duration)
}

func invoke(ctx context.Context, invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)

	destinationName := inv.Destination(kmReader).Name(kmReader)
	destinationLocation := inv.Destination(kmReader).Location(kmReader)
	if strings.HasPrefix(destinationName, "_") {
		if destinationLocation != "anywhere" {
			return nil, fmt.Errorf("capability '%s' location must be 'anywhere'", destinationName)
		}
		return invokeCapability(kmReader, inv)
	}

	if destinationLocation != "anywhere" {
		found := false
		for _, managedLocation := range managedLocations {
			if destinationLocation == managedLocation {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("invocationTargetLocation=%s; managedLocations=%v, rerouting\n",
				destinationLocation, managedLocations)
			return reroute(invBytes)
		}
	}

	moduleName := destinationName
	folderName := strings.Split(moduleName, "-")[0]
	module, err := getModule(moduleName)
	check(err)
	instance, err := module.Instantiate(ctx)
	defer instance.Close(ctx)
	check(err)

	// fmt.Printf("instance memory bytes before invoking: %d\n", instance.MemorySize(ctx))
	// fmt.Printf("invoking %s\n", folderName)
	result, err := instance.Invoke(ctx, folderName, invBytes)
	// fmt.Printf("instance memory bytes after invoking: %d\n", instance.MemorySize(ctx))
	return result, err
}

func invokeCapability(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	capabilityName := inv.Destination(kmReader).Name(kmReader)
	switch capabilityName {
	case "_wait":
		return wait.Handle(kmReader, inv)
	case "_async_wait":
		return wait.HandleAsync(kmReader, inv)
	case "_await_wait":
		return wait.HandleAwait(kmReader, inv)
	}
	return nil, fmt.Errorf("capability %s not implemented", capabilityName)
}

func reroute(invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)

	destinationLocation := inv.Destination(kmReader).Location(kmReader)
	switch destinationLocation {
	case "us-west1":
		return post("http://localhost:8081", invBytes)
	case "us-east1":
		return post("http://localhost:8082", invBytes)
	}
	return nil, fmt.Errorf("unsupported location: %s", destinationLocation)
}

func post(url string, invBytes []byte) ([]byte, error) {
	res, err := http.Post(url, "application/x-binary", bytes.NewBuffer(invBytes))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s", string(result))
	}
	return result, err
}

func host(ctx context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
	if binding != "" || namespace != "" || operation != "invoke" {
		return nil, fmt.Errorf("invalid host call, only 'invoke' is supported")
	}
	return invoke(ctx, payload)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
