package main

import (
	"fmt"
	"strings"

	wapc "github.com/wapc/wapc-guest-tinygo"
	"karmem.org/golang"

	waaskm "waas/km"
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
		"hello": helloWrapper,
	})
	fmt.Println("hello.main finished")
}

func helloWrapper(invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)
	location := inv.Destination(kmReader).Location(kmReader)

	for _, managedScope := range managedScopes {
		if location == managedScope {
			result, err := hello(string(inv.Payload(kmReader)))
			if err != nil {
				return nil, err
			} else {
				return []byte(result), nil
			}
		}
	}
	return nil, fmt.Errorf("managedScopes: %v; but found: %s", managedScopes, location)
}

func hello(name string) (string, error) {
	counter += 1
	fmt.Printf("hello with managedScopes %v called, counter = %d\n", managedScopes, counter)
	_ = make([]byte, 100)
	capitalized, err := invokeCapitalize(name)
	if err != nil {
		return "", err
	}
	// Format the message.
	msg := fmt.Sprintf("Hello, %s", capitalized)
	return msg, nil
}

func invokeCapitalize(str string) (string, error) {
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "hello",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     "capitalize",
			Location: "anywhere",
		},
		Payload:  []byte(str),
		Metadata: []waaskm.Metadata{},
	}
	kmWriter.Reset()
	_, err := inv.WriteAsRoot(kmWriter)
	if err != nil {
		return "", err
	}
	invBytes := kmWriter.Bytes()

	resultBytes, err := wapc.HostCall("", "", "invoke", invBytes)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
