package cmds

import (
	"bytes"
	"encoding/binary"

	"github.com/mraerino/atem-go/models"
)

type statusStringCommand string

func (s *statusStringCommand) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 44)
	copy(pl, []byte(*s))
	return pl, nil
}

func (s *statusStringCommand) UnmarshalBinary(data []byte) error {
	*s = statusStringCommand(decodeString(data))
	return nil
}

func (s *statusStringCommand) Value() string {
	return string(*s)
}

type PinCmd struct {
	statusStringCommand
}

func (PinCmd) Slug() string {
	return "_pin"
}

type WarnCmd struct {
	statusStringCommand
}

func (WarnCmd) Slug() string {
	return "Warn"
}

type TopCmd models.Topology

func (TopCmd) Slug() string {
	return "_top"
}

func (t *TopCmd) MarshalBinary() ([]byte, error) {
	fields := []uint8{
		t.MEs, t.Sources, t.ColorGenerators,
		t.AUXBusses, t.DownstreamKeyers, t.Stingers,
		t.DVEs, t.SuperSources,
		1, 0, 1, 0, // unknown
	}
	if t.SDOutput {
		fields[9] = 1
	}

	pl := make([]byte, 0, 12)
	buf := bytes.NewBuffer(pl)
	binary.Write(buf, binary.BigEndian, fields)
	return buf.Bytes(), nil
}

func (t *TopCmd) UnmarshalBinary(data []byte) error {
	buf := bytes.NewBuffer(data)
	var hasSDOut uint8
	fields := []interface{}{
		&t.MEs, &t.Sources, &t.ColorGenerators,
		&t.AUXBusses, &t.DownstreamKeyers, &t.Stingers,
		&t.DVEs, &t.SuperSources,
		new(uint8), &hasSDOut,
	}

	err := decode(buf, fields)
	if err != nil {
		return err
	}

	if hasSDOut > 0 {
		t.SDOutput = true
	}

	return nil
}

type MecCmd struct {
	ME     uint8
	Keyers uint8
}

func (MecCmd) Slug() string {
	return "_MeC"
}

func (m *MecCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = m.ME
	pl[1] = m.Keyers
	return pl, nil
}

func (m *MecCmd) UnmarshalBinary(data []byte) error {
	m.ME = data[0]
	m.Keyers = data[1]
	return nil
}

type MplCmd models.MediaPlayerConfig

func (MplCmd) Slug() string {
	return "_mpl"
}

func (m *MplCmd) MarshalBinary() ([]byte, error) {
	fields := []uint8{
		m.StillBanks, m.ClipBanks,
		0, 0, // unknown
	}
	pl := make([]byte, 0, 4)
	buf := bytes.NewBuffer(pl)
	binary.Write(buf, binary.BigEndian, fields)
	return buf.Bytes(), nil
}

func (m *MplCmd) UnmarshalBinary(data []byte) error {
	fields := []interface{}{
		&m.StillBanks, &m.ClipBanks,
	}
	buf := bytes.NewBuffer(data)
	return decode(buf, fields)
}

type MvcCmd uint8

func (MvcCmd) Slug() string {
	return "_MvC"
}

func (m *MvcCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = uint8(*m)
	return pl, nil
}

func (m *MvcCmd) UnmarshalBinary(data []byte) error {
	*m = MvcCmd(data[0])
	return nil
}

type SscCmd uint8

func (SscCmd) Slug() string {
	return "_SSC"
}

func (m *SscCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = uint8(*m)
	return pl, nil
}

func (m *SscCmd) UnmarshalBinary(data []byte) error {
	*m = SscCmd(data[0])
	return nil
}

type TlcCmd uint8

func (TlcCmd) Slug() string {
	return "_TlC"
}

func (t *TlcCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 8)
	pl[4] = uint8(*t)
	return pl, nil
}

func (t *TlcCmd) UnmarshalBinary(data []byte) error {
	*t = TlcCmd(data[4])
	return nil
}

type MacCmd uint8

func (MacCmd) Slug() string {
	return "_MAC"
}

func (m *MacCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = uint8(*m)
	return pl, nil
}

func (m *MacCmd) UnmarshalBinary(data []byte) error {
	*m = MacCmd(data[0])
	return nil
}

type PowrCmd models.PowerStatus

func (PowrCmd) Slug() string {
	return "Powr"
}

func (p *PowrCmd) MarshalBinary() ([]byte, error) {
	var status uint8
	if p.Main {
		status |= 1
	}
	if p.Backup {
		status |= 2
	}

	pl := make([]byte, 4)
	pl[0] = status

	return pl, nil
}

func (p *PowrCmd) UnmarshalBinary(data []byte) error {
	status := data[0]
	if status&1 != 0 {
		p.Main = true
	}
	if status&2 != 0 {
		p.Backup = true
	}
	return nil
}

type VidmCmd models.VideoMode

func (VidmCmd) Slug() string {
	return "VidM"
}

func (m *VidmCmd) MarshalBinary() ([]byte, error) {
	pl := make([]byte, 4)
	pl[0] = uint8(*m)
	return pl, nil
}

func (m *VidmCmd) UnmarshalBinary(data []byte) error {
	*m = VidmCmd(data[0])
	return nil
}
