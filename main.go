package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	engine := wazero.Engine()
	ctx := context.Background()

	guest, err := os.ReadFile("services/hello/hello.wasm")
	check(err)

	module, err := engine.New(ctx, host, guest, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	check(err)
	defer module.Close(ctx)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		test(ctx, module)
		wg.Done()
	}()
	go func() {
		test(ctx, module)
		wg.Done()
	}()
	wg.Wait()
}

func test(ctx context.Context, module wapc.Module) {
	instance, err := module.Instantiate(ctx)
	check(err)
	defer instance.Close(ctx)

	result, err := instance.Invoke(ctx, "hello", []byte("john"))
	check(err)

	fmt.Println(string(result))
}

func host(ctx context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
	switch operation {
	case "capitalize":
		name := string(payload)
		name = strings.Title(name)
		return []byte(name), nil
	}
	return nil, fmt.Errorf("operation name not found: %s", operation)
}