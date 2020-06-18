package cmds

import (
	"encoding/binary"

	"github.com/mraerino/atem-go/models"
)

type TlinCmd map[int]models.TallyState

func (TlinCmd) Slug() string {
	return "TlIn"
}

func (t TlinCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, len(t)+2)
	binary.BigEndian.PutUint16(pl, uint16(len(t)))

	for idx, state := range t {
		pl[idx] = state.Bitmask()
	}

	return pl, nil
}

func (t TlinCmd) UnmarshalBinary(data []byte) error {
	num := binary.BigEndian.Uint16(data)

	for i := 0; i < int(num); i++ {
		t[i] = models.TallyState{
			Program: data[2+i]&1 != 0,
			Preview: data[2+i]&2 != 0,
		}
	}

	return nil
}

type TlsrCmd map[models.VideoSource]models.TallyState

func (TlsrCmd) Slug() string {
	return "TlSr"
}

func (t TlsrCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, len(t)*3+2)
	binary.BigEndian.PutUint16(pl, uint16(len(t)))

	var i int
	for idx, state := range t {
		start := 2 + i*3
		binary.BigEndian.PutUint16(pl[start:], uint16(idx))
		pl[start+2] = state.Bitmask()
		i++
	}

	return pl, nil
}

func (t TlsrCmd) UnmarshalBinary(data []byte) error {
	num := binary.BigEndian.Uint16(data)

	for i := 0; i < int(num); i++ {
		start := 2 + i*3
		src := models.VideoSource(binary.BigEndian.Uint16(data[start:]))
		t[src] = models.TallyState{
			Program: data[start+2]&1 != 0,
			Preview: data[start+2]&2 != 0,
		}
	}

	return nil
}
