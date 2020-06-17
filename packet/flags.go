package packet

import "github.com/sirupsen/logrus"

const (
	FlagNeedACK Flags = 1 << iota
	FlagInit
	FlagRetrans
	FlagHello
	FlagACK
)

type Flags uint8

func (f Flags) Has(mask Flags) bool {
	return uint8(f)&uint8(mask) != 0
}

func (f Flags) Debug() logrus.Fields {
	flags := map[string]Flags{
		"needACK":        FlagNeedACK,
		"init":           FlagInit,
		"retransmission": FlagRetrans,
		"hello":          FlagHello,
		"ack":            FlagACK,
	}
	fields := make(logrus.Fields)
	for name, val := range flags {
		fields[name] = f.Has(val)
	}
	return fields
}

func FlagsFrom(masks ...Flags) Flags {
	var flags Flags
	for _, mask := range masks {
		flags |= mask
	}
	return flags
}
