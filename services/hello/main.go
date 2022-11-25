package main

import (
	"fmt"
	"strings"

	wapc "github.com/wapc/wapc-guest-tinygo"
	"karmem.org/golang"

	"hello/waaskm"
)

var managedScopesFlag string
var managedScopes []string

var counter int32
var kmWriter *karmem.Writer = karmem.NewWriter(1024)

func main() {
	managedScopes = strings.Split(managedScopesFlag, ",")

	counter = 0
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"hello": hello,
	})
	fmt.Println("hello.main finished")
}

func hello(invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)
	location := inv.Destination(kmReader).Location(kmReader)
	
	for _, managedScope := range managedScopes {
		if location == managedScope {
			return helloInternal(inv.Payload(kmReader))
		}
	}
	return nil, fmt.Errorf("managedScopes: %v; but found: %s", location)
}

// hello will callback the host and return the payload
func helloInternal(payload []byte) ([]byte, error) {
	counter += 1
	fmt.Printf("hello with managedScopes %v called, counter = %d\n", managedScopes, counter)
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
	kmWriter.Reset()
	_, err := inv.WriteAsRoot(kmWriter);
	if err != nil {
		return nil, err
	}
	invBytes := kmWriter.Bytes()

	nameBytes, err := wapc.HostCall("", "", "invoke", invBytes)
	if err != nil {
		return nil, err
	}
	return nameBytes, nil
}