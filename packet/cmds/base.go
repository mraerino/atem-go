package cmds

import (
	"encoding/binary"

	"github.com/mraerino/atem-go/models"
)

type UnknownCommand struct {
	slug string
	data []byte
}

func NewUnknownCommand(slug string) *UnknownCommand {
	return &UnknownCommand{slug: slug}
}

func (u *UnknownCommand) Slug() string {
	return u.slug
}

func (u *UnknownCommand) MarshalBinary() ([]byte, error) {
	return u.data, nil
}

func (u *UnknownCommand) UnmarshalBinary(data []byte) error {
	u.data = data[:]
	return nil
}

// VerCmd represents the _ver command
type VerCmd struct {
	Major uint16
	Minor uint16
}

func (VerCmd) Slug() string {
	return "_ver"
}

func (c *VerCmd) MarshalBinary() ([]byte, error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint16(buf, c.Major)
	binary.BigEndian.PutUint16(buf[2:], c.Minor)
	return buf, nil
}

func (c *VerCmd) UnmarshalBinary(data []byte) error {
	c.Major = binary.BigEndian.Uint16(data)
	c.Minor = binary.BigEndian.Uint16(data[2:])
	return nil
}

type IncmCmd struct{}

func (*IncmCmd) Slug() string {
	return "InCm"
}

func (*IncmCmd) MarshalBinary() ([]byte, error) {
	return make([]byte, 4), nil // todo: handle states
}

func (*IncmCmd) UnmarshalBinary(data []byte) error {
	return nil
}

type TimeCmd models.Timecode

func (TimeCmd) Slug() string {
	return "Time"
}

func (t *TimeCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 8)
	pl[0] = uint8(t.Hour)
	pl[1] = uint8(t.Minute)
	pl[2] = uint8(t.Second)
	pl[3] = uint8(t.Frame)
	return pl, nil
}

func (t *TimeCmd) UnmarshalBinary(data []byte) error {
	t.Hour = int(data[0])
	t.Minute = int(data[1])
	t.Second = int(data[2])
	t.Frame = int(data[3])
	return nil
}

type TiRqCmd struct{}

func (TiRqCmd) Slug() string {
	return "TiRq"
}

func (TiRqCmd) MarshalBinary() ([]byte, error) {
	return make([]byte, 0), nil
}

func (TiRqCmd) UnmarshalBinary(_ []byte) error {
	return nil
}
