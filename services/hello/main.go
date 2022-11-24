package main

import (
	"fmt"

	wapc "github.com/wapc/wapc-guest-tinygo"
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
	// Make a host call to capitalize the name.
	nameBytes, err := wapc.HostCall("", "", "capitalize", payload)
	if err != nil {
		return nil, err
	}
	// Format the message.
	msg := "Hello, " + string(nameBytes)
	return []byte(msg), nil
}
