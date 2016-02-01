package mbproto

import "testing"
import (
	"github.com/golang/protobuf/proto"
	"bytes"
)

func TestEncode(t *testing.T) {
	rm := RawMessage{
		IV: bytes.Repeat([]byte("O"), 16),
		Nonce: bytes.Repeat([]byte("O"), 16),
		Payload: bytes.Repeat([]byte("O"), 1400),
	}
	data, err := proto.Marshal(&rm)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.Buffer{}
	t.Logf("len %d", len(data))
	for i := 0; i < 10; i++ {
		buf.Write(data)
	}
}