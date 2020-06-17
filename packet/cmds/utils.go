package cmds

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

// decode decodes network packets
func decode(buf *bytes.Buffer, fields []interface{}) error {
	var err error
	for _, field := range fields {
		err = binary.Read(buf, binary.BigEndian, field)
		if err != nil {
			return errors.Wrap(err, "Unable to read from buffer")
		}
	}
	return nil
}

func decodeString(data []byte) string {
	end := len(data)
	for i, c := range data {
		if c == 0 {
			end = i
			break
		}
	}
	return string(data[:end])
}
