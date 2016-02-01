package mbproto

import (
	"io"
	"encoding/binary"
)

func ReadRawMessage(r io.Reader) (int, error) {
	var (
		len uint16
		readed int
	)
	err := binary.Read(r, binary.BigEndian, &len)
	if err != nil {
		return readed, err
	}
	return 0, nil
}