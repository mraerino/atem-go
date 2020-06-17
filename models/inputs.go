package models

type InputProperties struct {
	SourceIndex VideoSource
	LongName    string
	ShortName   string

	ExternalPortType ExternalPortType
	PortType         PortType
}

// Generate String() methods
//go:generate go run golang.org/x/tools/cmd/stringer -type=ExternalPortType,PortType -linecomment -output=inputs_string.go

type ExternalPortType uint8

const (
	ExternalPortTypeInternal  ExternalPortType = iota // Internal
	ExternalPortTypeSDI                               // SDI
	ExternalPortTypeHDMI                              // HDMI
	ExternalPortTypeComposite                         // Composite
	ExternalPortTypeComponent                         // Component
	ExternalPortTypeSVideo                            // SVideo
)

type PortType uint8

const (
	PortTypeExternal        PortType = iota // External
	PortTypeBlack                           // Black
	PortTypeColorBars                       // Color Bars
	PortTypeColorGenerator                  // Color Generator
	PortTypeMediaPlayerFill                 // Media Player Fill
	PortTypeMediaPlayerKey                  // Media Player Key
	PortTypeSuperSource                     // SuperSource
)

const (
	PortTypeMEOutput PortType = 128 + iota // ME Output
	PortTypeAuxilary                       // Auxilary
	PortTypeMask                           // Mask
)
