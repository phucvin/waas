package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"karmem.org/golang"
	waaskm "waas/km"
)

var kmWriterPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func main() {
	test()
}

func test() {
	kmWriter := kmWriterPool.Get().(*karmem.Writer)
	defer kmWriterPool.Put(kmWriter)
	defer kmWriter.Reset()
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "test02",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     "ping",
			Location: "anywhere",
		},
		Payload:  []byte{99},
		Metadata: []waaskm.Metadata{},
	}
	_, err := inv.WriteAsRoot(kmWriter)
	check(err)
	invBytes := kmWriter.Bytes()

	res, err := http.Post("http://localhost:8081", "application/x-binary", bytes.NewBuffer(invBytes))
	check(err)
	result, err := ioutil.ReadAll(res.Body)
	check(err)
	fmt.Printf("%s: %v\n", res.Status, result)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
