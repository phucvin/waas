package main

import (
	"context"
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
)

var engine wapc.Engine
var moduleMapMutex = &sync.RWMutex{}
var moduleCtx context.Context = context.Background()
var moduleMap map[string]*wapc.Pool

var kmWriterPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func resetModules() {
	func() {
		moduleMapMutex.RLock()
		defer moduleMapMutex.RUnlock()
		for _, module := range moduleMap {
			module.Close(moduleCtx)
		}
	}()

	moduleMapMutex.Lock()
	defer moduleMapMutex.Unlock()
	moduleMap = make(map[string]*wapc.Pool)
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

	pool, err := wapc.NewPool(moduleCtx, module, 2, func(instance wapc.Instance) error {
		return nil
	})

	moduleMapMutex.Lock()
	defer moduleMapMutex.Unlock()
	moduleMap[name] = pool
}

func getModule(name string) (*wapc.Pool, error) {
	moduleMapMutex.RLock()
	defer moduleMapMutex.RUnlock()
	if module, exist := moduleMap[name]; exist {
		return module, nil
	}
	return nil, fmt.Errorf("module not loaded or not found: %s", name)
}

func main() {
	engine = wazero.Engine()
	resetModules()

	fmt.Println("loading modules")
	loadModule("hello-us-west1")
	loadModule("hello-us-east1")
	loadModule("capitalize")
	fmt.Println("3 modules loaded")

	handler := http.HandlerFunc(handleHTTP)
	fmt.Println("Listening...")
	http.ListenAndServe(":8080", handler)

	resetModules() // close loaded modules
}

func handleHTTP(w http.ResponseWriter, req *http.Request) {
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
}

func invoke(ctx context.Context, invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)
	moduleName := inv.Destination(kmReader).Name(kmReader)
	folderName := strings.Split(moduleName, "-")[0]

	modulePool, err := getModule(moduleName)
	check(err)
	instance, err := modulePool.Get(5 * time.Millisecond)
	check(err)
	defer modulePool.Return(instance)

	return instance.Invoke(ctx, folderName, invBytes)
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
