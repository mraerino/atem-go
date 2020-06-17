package cmds

import (
	"encoding/binary"

	"github.com/mraerino/atem-go/models"
)

type InprCmd models.InputProperties

func (InprCmd) Slug() string {
	return "InPr"
}

func (i *InprCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 36)
	pl[0] = byte(i.SourceIndex)
	copy(pl[2:22], []byte(i.LongName))
	copy(pl[22:26], []byte(i.ShortName))

	// todo: figure out missing bits

	pl[29] = byte(i.ExternalPortType)
	pl[30] = byte(i.PortType)

	return pl, nil
}

func (i *InprCmd) UnmarshalBinary(data []byte) error {
	src := binary.BigEndian.Uint16(data)
	i.SourceIndex = models.VideoSource(src)

	i.LongName = decodeString(data[2:22])
	i.ShortName = decodeString(data[22:26])

	// todo: figure out rest of the bytes

	i.ExternalPortType = models.ExternalPortType(data[29])
	i.PortType = models.PortType(data[30])

	return nil
}
