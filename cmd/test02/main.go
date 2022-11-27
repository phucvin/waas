package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"karmem.org/golang"
	waaskm "waas/km"
)

var kmWriterPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func main() {
	n := flag.Int("n", 1, "")
	flag.Parse()
	start := time.Now()
	test(*n)
	duration := time.Since(start)
	fmt.Printf("test took %v\n", duration)
}

func test(n int) {
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
		Payload:  []byte{byte(n)},
		Metadata: []waaskm.Metadata{},
	}
	_, err := inv.WriteAsRoot(kmWriter)
	check(err)
	invBytes := kmWriter.Bytes()

	res, err := http.Post("http://localhost:8082", "application/x-binary", bytes.NewBuffer(invBytes))
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
