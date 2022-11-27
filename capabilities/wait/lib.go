package wait

import (
	"encoding/binary"
	"fmt"
	"time"
	"sync"
	"sync/atomic"
	
	"karmem.org/golang"
	waaskm "waas/km"
)

var tokenMap *sync.Map
var nextToken uint64

func Reset() {
	tokenMap = &sync.Map{}
	nextToken = 0
}

func Handle(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	payload := inv.Payload(kmReader)
	if (len(payload) != 4) {
		return nil, fmt.Errorf("invalid payload size for _wait")
	}
	milliseconds := binary.LittleEndian.Uint32(payload)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
	return []byte{}, nil
}

func HandleAsync(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	payload := inv.Payload(kmReader)
	if (len(payload) != 4) {
		return nil, fmt.Errorf("invalid payload size for _wait")
	}
	milliseconds := binary.LittleEndian.Uint32(payload)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		wg.Done()
	}()
	token := atomic.AddUint64(&nextToken, 1)
	if token % 5000 == 0 {
		fmt.Printf("wait token reached %d\n", token)
	}
	tokenMap.Store(token, &wg)
	tokenBytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(tokenBytes, token)
	return tokenBytes, nil
}

func HandleAwait(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	token := binary.LittleEndian.Uint64(inv.Payload(kmReader))
	if token % 5000 == 0 {
		fmt.Printf("awaiting token %d\n", token)
	}
	wg, ok := tokenMap.LoadAndDelete(token)
	if !ok {
		return nil, fmt.Errorf("token not found while handling _await_wait: %v", token)
	}
	wg.(*sync.WaitGroup).Wait()
	return []byte{}, nil
}