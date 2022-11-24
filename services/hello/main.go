package main

import (
	"fmt"

	wapc "github.com/wapc/wapc-guest-tinygo"
	"karmem.org/golang"

	"hello/waaskm"
)

var counter int32

func main() {
	counter = 0
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"hello": hello,
	})
	fmt.Println("hello.main finished")
}

// hello will callback the host and return the payload
func hello(payload []byte) ([]byte, error) {
	counter += 1
	fmt.Printf("hello called, counter = %d\n", counter)
	_ = make([]byte, 100)
	nameBytes, err := invokeCapitalize(payload)
	if err != nil {
		return nil, err
	}
	// Format the message.
	msg := "Hello, " + string(nameBytes)
	return []byte(msg), nil
}

func invokeCapitalize(payload []byte) ([]byte, error) {
	writer := karmem.NewWriter(1024)
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name: "hello",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name: "capitalize",
			Location: "global",
		},
		Payload: payload,
		Metadata: []waaskm.Metadata{},
	}
	_, err := inv.WriteAsRoot(writer);
	if err != nil {
		return nil, err
	}
	invBytes := writer.Bytes()

	nameBytes, err := wapc.HostCall("", "", "invoke", invBytes)
	if err != nil {
		return nil, err
	}
	return nameBytes, nil
}