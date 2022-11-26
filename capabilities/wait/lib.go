package wait

import (
	"encoding/binary"
	"fmt"
	"time"

	
	"karmem.org/golang"
	waaskm "waas/km"
)

func Handle(kmReader *karmem.Reader, inv *waaskm.InvocationViewer) ([]byte, error) {
	payload := inv.Payload(kmReader)
	if (len(payload) != 4) {
		return nil, fmt.Errorf("invalid payload size for _wait")
	}
	milliseconds := binary.LittleEndian.Uint32(payload)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
	return []byte{}, nil
}