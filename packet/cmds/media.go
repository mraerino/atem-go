package cmds

import (
	"encoding/binary"
	"errors"

	"github.com/mraerino/atem-go/models"
)

type MpceCmd struct {
	PlayerIndex uint8
	Type        models.MediaPlayerType
	StillIndex  int
	ClipIndex   int
}

func (MpceCmd) Slug() string {
	return "MPCE"
}

func (m *MpceCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = m.PlayerIndex
	pl[1] = uint8(m.Type)
	pl[2] = uint8(m.StillIndex)
	pl[3] = uint8(m.ClipIndex)
	return pl, nil
}

func (m *MpceCmd) UnmarshalBinary(data []byte) error {
	m.PlayerIndex = data[0]
	m.Type = models.MediaPlayerType(data[1])
	m.StillIndex = int(data[2])
	m.ClipIndex = int(data[3])
	return nil
}

type MpfeCmd struct {
	Index    uint16
	Used     bool
	Hash     []byte
	Filename string
}

func (MpfeCmd) Slug() string {
	return "MPfe"
}

func (m *MpfeCmd) MarshalBinary() ([]byte, error) {
	filenameBytes := []byte(m.Filename)
	pl := make([]byte, 24+len(filenameBytes))

	pl[0] = 0 // still type
	binary.BigEndian.PutUint16(pl[2:], m.Index)
	if m.Used {
		pl[4] = 1
	}

	copy(pl[5:], []byte(m.Hash)[:16])

	pl[23] = uint8(len(filenameBytes))
	copy(pl[24:], filenameBytes)

	return pl, nil
}

func (m *MpfeCmd) UnmarshalBinary(data []byte) error {
	if data[0] != 0 {
		return errors.New("Unknown file type")
	}

	m.Index = binary.BigEndian.Uint16(data[2:])
	m.Used = data[4] == 1
	m.Hash = data[5:21]

	filenameLen := data[23]
	m.Filename = string(data[24 : 24+filenameLen])

	return nil
}
