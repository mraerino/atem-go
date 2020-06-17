package packet

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeserializeHeader(t *testing.T) {
	log := logrus.New()

	cases := map[string]struct {
		pl             []byte
		expectedFlags  Flags
		expectedLength uint16
	}{
		"Init": {
			pl: []byte{
				0x10,       // Flags
				0x14,       // Length
				0x1e, 0xbd, // Session ID
				0x00, 0x00, // Acked Seq Num
				0x00, 0x00, // Unknown
				0x00, 0xa8, // ???
				0x00, 0x00, // Seq Num

				// Payload
				0x01, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			expectedFlags:  FlagsFrom(FlagInit),
			expectedLength: 20,
		},
		"ACK": {
			pl: []byte{
				0x80, // Flags
				0x0c, // Length
				0x50, 0x5e,
				0x00, 0x00,
				0x00, 0x00,
				0x00, 0x6d,
				0x00, 0x00,
			},
			expectedFlags:  FlagsFrom(FlagACK),
			expectedLength: 12,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			log := log.WithField("test", t.Name())
			buf := bytes.NewBuffer(test.pl)

			msg, err := Deserialize(log, buf)
			require.NoError(t, err)

			assert.Equal(t, test.expectedFlags, msg.Flags)
			assert.EqualValues(t, test.expectedLength, msg.Length)
		})
	}
}
