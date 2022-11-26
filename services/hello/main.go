package main

import (
	"encoding/binary"
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
	// fmt.Println("hello.main finished")
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

	waitTokens := make([][]byte, 100000)
	for i := 0; i < len(waitTokens); i += 1 {
		waitToken, err := asyncWait(100 /* milliseconds */)
		if err != nil {
			return "", err
		}
		waitTokens[i] = waitToken
	}

	capitalized, err := invokeCapitalize(name)
	if err != nil {
		return "", err
	}
	// Format the message.
	msg := fmt.Sprintf("Hello, %s", capitalized)

	for i := 0; i < len(waitTokens); i += 1 {
		_, err := awaitWait(waitTokens[i])
		if err != nil {
			return "", err
		}
	}

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

func invokeWait(milliseconds uint32) error {
	millisecondsBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(millisecondsBytes, milliseconds)
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "hello",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     "_wait",
			Location: "anywhere",
		},
		Payload:  millisecondsBytes,
		Metadata: []waaskm.Metadata{},
	}
	kmWriter.Reset()
	_, err := inv.WriteAsRoot(kmWriter)
	if err != nil {
		return  err
	}
	invBytes := kmWriter.Bytes()

	_, err = wapc.HostCall("", "", "invoke", invBytes)
	return err
}

func asyncWait(milliseconds uint32) ([]byte, error) {
	millisecondsBytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(millisecondsBytes, milliseconds)
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "hello",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     "_async_wait",
			Location: "anywhere",
		},
		Payload:  millisecondsBytes,
		Metadata: []waaskm.Metadata{},
	}
	kmWriter.Reset()
	_, err := inv.WriteAsRoot(kmWriter)
	if err != nil {
		return  nil, err
	}
	invBytes := kmWriter.Bytes()

	tokenBytes, err := wapc.HostCall("", "", "invoke", invBytes)
	return tokenBytes, err
}

func awaitWait(token []byte) ([]byte, error) {
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "hello",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     "_await_wait",
			Location: "anywhere",
		},
		Payload:  token,
		Metadata: []waaskm.Metadata{},
	}
	kmWriter.Reset()
	_, err := inv.WriteAsRoot(kmWriter)
	if err != nil {
		return  nil, err
	}
	invBytes := kmWriter.Bytes()

	return wapc.HostCall("", "", "invoke", invBytes)
}