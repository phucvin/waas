package main

import (
	"fmt"
	"strings"

	wapc "github.com/wapc/wapc-guest-tinygo"
)

var counter int32

func main() {
	counter = 0
	// Register echo and fail functions
	wapc.RegisterFunctions(wapc.Functions{
		"capitalize": capitalize,
	})
}

// captialize will change the string
func capitalize(payload []byte) ([]byte, error) {
	counter += 1
	fmt.Printf("capitalize called, counter = %d\n", counter)
	_ = make([]byte, 100)
	return []byte(strings.Title(string(payload))), nil
}
