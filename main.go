package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
	"karmem.org/golang"

	waaskm "waas/km"
)

var engine wapc.Engine
var moduleMapMutex = &sync.RWMutex{}
var moduleCtx context.Context = context.Background()
var moduleMap map[string]wapc.Module

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
	moduleMap = make(map[string]wapc.Module)
}

func loadModule(name string) {
	wasm, err := os.ReadFile(fmt.Sprintf("services/%s/%s.wasm", name, name))
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
	engine = wazero.Engine()
	resetModules()

	fmt.Println("loading modules")
	loadModule("hello")
	loadModule("capitalize")
	fmt.Println("2 modules loaded")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		test()
		wg.Done()
	}()
	go func() {
		test()
		wg.Done()
	}()
	wg.Wait()

	resetModules() // close loaded modules
}

func test() {
	kmWriter := kmWriterPool.Get().(*karmem.Writer)
	defer kmWriterPool.Put(kmWriter)
	defer kmWriter.Reset()
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name: "test",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name: "hello",
			Location: "global",
		},
		Payload: []byte("bob"),
		Metadata: []waaskm.Metadata{},
	}
	_, err := inv.WriteAsRoot(kmWriter);
	check(err)
	invBytes := kmWriter.Bytes()

	ctx := moduleCtx
	result, err := invoke(ctx, invBytes)
	check(err)
	fmt.Println(string(result))
}

func invoke(ctx context.Context, invBytes []byte) ([]byte, error) {
	reader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(reader, 0)
	moduleName := inv.Destination(reader).Name(reader)
	
	module, err := getModule(moduleName)
	check(err)
	instance, err := module.Instantiate(ctx)
	check(err)
	fmt.Printf("instance initialized: %s\n", moduleName)
	defer instance.Close(ctx)

	return instance.Invoke(ctx, moduleName, inv.Payload(reader))
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
