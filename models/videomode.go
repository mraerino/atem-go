package models

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer -type=VideoMode -linecomment
type VideoMode uint8

const (
	VideoMode525i59_94NTSC     VideoMode = iota // 525i59.94 NTSC
	VideoMode625i50PAL                          // 625i50 PAL
	VideoMode525i59_94NTSC16_9                  // 525i59.94 NTSC 16:9
	VideoMode625i50PAL16_9                      // 625i50 PAL 16:9
	VideoMode720p50                             // 720p50
	VideoMode720p59_94                          // 720p59.94
	VideoMode1080i50                            // 1080i50
	VideoMode1080i59_94                         // 1080i59.94
	VideoMode1080p23_98                         // 1080p23.98
	VideoMode1080p24                            // 1080p24
	VideoMode1080p25                            // 1080p25
	VideoMode1080p29_97                         // 1080p29.97
	VideoMode1080p50                            // 1080p50
	VideoMode1080p59_94                         // 1080p59.94
	VideoMode2160p23_98                         // 2160p23.98
	VideoMode2160p24                            // 2160p24
	VideoMode2160p25                            // 2160p25
	VideoMode2160p29_97                         // 2160p29.97
)

// new modes
const (
	VideoMode1080p30 VideoMode = 26 + iota // 1080p30
	VideoMode1080p60                       // 1080p60
)

// todo: remove
func (m VideoMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, m.String())), nil
}
