package models

//go:generate go run golang.org/x/tools/cmd/stringer -type=AudioSource -linecomment
type AudioSource uint16

const (
	AudioSourceInput1  AudioSource = 1 + iota // Input 1
	AudioSourceInput2                         // Input 2
	AudioSourceInput3                         // Input 3
	AudioSourceInput4                         // Input 4
	AudioSourceInput5                         // Input 5
	AudioSourceInput6                         // Input 6
	AudioSourceInput7                         // Input 7
	AudioSourceInput8                         // Input 8
	AudioSourceInput9                         // Input 9
	AudioSourceInput10                        // Input 10
	AudioSourceInput11                        // Input 11
	AudioSourceInput12                        // Input 12
	AudioSourceInput13                        // Input 13
	AudioSourceInput14                        // Input 14
	AudioSourceInput15                        // Input 15
	AudioSourceInput16                        // Input 16
	AudioSourceInput17                        // Input 17
	AudioSourceInput18                        // Input 18
	AudioSourceInput19                        // Input 19
	AudioSourceInput20                        // Input 20
)
const (
	AudioSourceXLR    AudioSource = 1001 // XLR
	AudioSourceAESEBU AudioSource = 1101 // AES/EBU
	AudioSourceRCA    AudioSource = 1201 // RCA
	AudioSourceMic1   AudioSource = 1301 // Mic 1
	AudioSourceMic2   AudioSource = 1302 // Mic 2
	AudioSourceMP1    AudioSource = 2001 // MP1
	AudioSourceMP2    AudioSource = 2002 // MP2
)
