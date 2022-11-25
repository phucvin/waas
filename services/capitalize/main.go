package main

import (
	"fmt"
	"strings"

	wapc "github.com/wapc/wapc-guest-tinygo"
	"karmem.org/golang"

	"capitalize/waaskm"
)

var counter int32

func main() {
	counter = 0
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"capitalize": capitalize,
	})
	fmt.Println("capitalize.main finished")
}

func capitalize(invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)
	result, err := capitalizeInternal(string(inv.Payload(kmReader)))
	if err != nil {
		return nil, err
	} else {
		return []byte(result), nil
	}
}

// captialize will change the string
func capitalizeInternal(str string) (string, error) {
	counter += 1
	fmt.Printf("capitalize called, counter = %d\n", counter)
	return strings.Title(str), nil
}
