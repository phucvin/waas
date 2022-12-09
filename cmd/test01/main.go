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
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 1; i++ {
		go func() {
			test("hello-us-west1", "us-west1")
			wg.Done()
		}()
		go func() {
			test("hello-us-east1", "us-east1")
			wg.Done()
		}()
	}
	wg.Wait()
}

func test(name, location string) {
	kmWriter := kmWriterPool.Get().(*karmem.Writer)
	defer kmWriterPool.Put(kmWriter)
	defer kmWriter.Reset()
	inv := waaskm.Invocation{
		Source: waaskm.Source{
			Name:     "test01",
			Location: "global",
		},
		Destination: waaskm.Destination{
			Name:     name,
			Location: location,
		},
		Payload:  []byte("bob"),
		Metadata: []waaskm.Metadata{},
	}
	_, err := inv.WriteAsRoot(kmWriter)
	check(err)
	invBytes := kmWriter.Bytes()

	res, err := http.Post("http://localhost:8082", "application/x-binary", bytes.NewBuffer(invBytes))
	check(err)
	result, err := ioutil.ReadAll(res.Body)
	check(err)
	fmt.Printf("%s: %s\n", res.Status, string(result))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
