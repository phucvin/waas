package main

import (
	"fmt"
	"strings"

	wapc "github.com/wapc/wapc-guest-tinygo"
	"karmem.org/golang"

	waaskm "waas/km"
)

var counter int32

func main() {
	counter = 0
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"capitalize": capitalizeWrapper,
	})
	fmt.Println("capitalize.main finished")
}

func capitalizeWrapper(invBytes []byte) ([]byte, error) {
	kmReader := karmem.NewReader(invBytes)
	inv := waaskm.NewInvocationViewer(kmReader, 0)
	result, err := capitalize(string(inv.Payload(kmReader)))
	if err != nil {
		return nil, err
	} else {
		return []byte(result), nil
	}
}

func capitalize(str string) (string, error) {
	counter += 1
	fmt.Printf("capitalize called, counter = %d\n", counter)
	return strings.Title(str), nil
}
