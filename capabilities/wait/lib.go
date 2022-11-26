package wait

import (
	"encoding/binary"
	"fmt"
	"time"
	"sync"
	
	"karmem.org/golang"
	waaskm "waas/km"
)

var tokenMap map[byte]*sync.WaitGroup
var nextToken byte

func Reset() {
	tokenMap = make(map[byte]*sync.WaitGroup)
	nextToken = 1
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
	token := nextToken
	nextToken = nextToken + 1
	tokenMap[token] = &wg
	return []byte{token}, nil
}

func HandleAwait(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	payload := inv.Payload(kmReader)
	if len(payload) != 1 {
		return nil, fmt.Errorf("invalid token sent to _await_wait: %v", payload)
	}
	token := payload[0]
	wg, ok := tokenMap[token]
	if !ok {
		return nil, fmt.Errorf("token not found while handling _await_wait: %v", token)
	}
	wg.Wait()
	delete(tokenMap, token)
	return []byte{}, nil
}