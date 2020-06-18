package packet

import (
	"encoding"
	"encoding/binary"

	"github.com/mraerino/atem-go/packet/cmds"
)

type Command interface {
	Slug() string
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// UnmarshalCommand will deserialize a command from bytes (excluding length field)
func UnmarshalCommand(data []byte) (Command, error) {
	var cmd Command

	// 2 unknown bytes preceding, ignoring right now
	slug := string(data[2:6])
	switch slug {
	case "InCm":
		cmd = new(cmds.IncmCmd)
	case "_ver":
		cmd = new(cmds.VerCmd)
	case "_pin":
		cmd = new(cmds.PinCmd)
	case "Warn":
		cmd = new(cmds.WarnCmd)
	case "_top":
		cmd = new(cmds.TopCmd)
	case "_MeC":
		cmd = new(cmds.MecCmd)
	case "_mpl":
		cmd = new(cmds.MplCmd)
	case "_MvC":
		cmd = new(cmds.MvcCmd)
	case "_SSC":
		cmd = new(cmds.SscCmd)
	case "_TlC":
		cmd = new(cmds.TlcCmd)
	case "_MAC":
		cmd = new(cmds.MacCmd)
	case "Powr":
		cmd = new(cmds.PowrCmd)
	case "VidM":
		cmd = new(cmds.VidmCmd)
	case "InPr":
		cmd = new(cmds.InprCmd)
	case "PrgI":
		cmd = new(cmds.PrgiCmd)
	case "PrvI":
		cmd = new(cmds.PrviCmd)
	case "AuxS":
		cmd = new(cmds.AuxsCmd)
	case "MPCE":
		cmd = new(cmds.MpceCmd)
	case "MPfe":
		cmd = new(cmds.MpfeCmd)
	case "TlIn":
		cmd = make(cmds.TlinCmd)
	case "TlSr":
		cmd = make(cmds.TlsrCmd)
	case "Time":
		cmd = new(cmds.TimeCmd)
	default:
		// unknown command (yet)
		cmd = cmds.NewUnknownCommand(slug)
	}

	return cmd, cmd.UnmarshalBinary(data[6:])
}

// MarshalCommand will serialize a command including the length field
func MarshalCommand(cmd Command) ([]byte, error) {
	// payload
	pl, err := cmd.MarshalBinary()
	if err != nil {
		return nil, err
	}

	// length
	out := make([]byte, 2)
	binary.BigEndian.PutUint16(out, uint16(len(pl))+8) // 8 bytes "header"

	// mystery 2 bytes
	out = append(out, []byte{0, 0}...)

	// slug
	out = append(out, []byte(cmd.Slug())[:4]...) // force to 4 bytes

	out = append(out, pl...)
	return out, nil
}
