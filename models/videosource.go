package models

//go:generate go run golang.org/x/tools/cmd/stringer -type=VideoSource -linecomment
type VideoSource uint16

const (
	VideoSourceBlack   VideoSource = iota // Black
	VideoSourceInput1                     // Input 1
	VideoSourceInput2                     // Input 2
	VideoSourceInput3                     // Input 3
	VideoSourceInput4                     // Input 4
	VideoSourceInput5                     // Input 5
	VideoSourceInput6                     // Input 6
	VideoSourceInput7                     // Input 7
	VideoSourceInput8                     // Input 8
	VideoSourceInput9                     // Input 9
	VideoSourceInput10                    // Input 10
	VideoSourceInput11                    // Input 11
	VideoSourceInput12                    // Input 12
	VideoSourceInput13                    // Input 13
	VideoSourceInput14                    // Input 14
	VideoSourceInput15                    // Input 15
	VideoSourceInput16                    // Input 16
	VideoSourceInput17                    // Input 17
	VideoSourceInput18                    // Input 18
	VideoSourceInput19                    // Input 19
	VideoSourceInput20                    // Input 20
)

// Special sources
const (
	VideoSourceColorBars       VideoSource = 1000  // ColorBars
	VideoSourceColor1          VideoSource = 2001  // Color 1
	VideoSourceColor2          VideoSource = 2002  // Color 2
	VideoSourceMediaPlayer1    VideoSource = 3010  // MediaPlayer 1
	VideoSourceMediaPlayer1Key VideoSource = 3011  // MediaPlayer 1 Key
	VideoSourceMediaPlayer2    VideoSource = 3020  // MediaPlayer 2
	VideoSourceMediaPlayer2Key VideoSource = 3021  // MediaPlayer 2 Key
	VideoSourceKey1Mask        VideoSource = 4010  // Key 1 Mask
	VideoSourceKey2Mask        VideoSource = 4020  // Key 2 Mask
	VideoSourceKey3Mask        VideoSource = 4030  // Key 3 Mask
	VideoSourceKey4Mask        VideoSource = 4040  // Key 4 Mask
	VideoSourceDSK1Mask        VideoSource = 5010  // DSK 1 Mask
	VideoSourceDSK2Mask        VideoSource = 5020  // DSK 2 Mask
	VideoSourceSuperSource     VideoSource = 6000  // SuperSource
	VideoSourceCleanFeed1      VideoSource = 7001  // CleanFeed 1
	VideoSourceCleanFeed2      VideoSource = 7002  // CleanFeed 2
	VideoSourceAuxilary1       VideoSource = 8001  // Auxilary 1
	VideoSourceAuxilary2       VideoSource = 8002  // Auxilary 2
	VideoSourceAuxilary3       VideoSource = 8003  // Auxilary 3
	VideoSourceAuxilary4       VideoSource = 8004  // Auxilary 4
	VideoSourceAuxilary5       VideoSource = 8005  // Auxilary 5
	VideoSourceAuxilary6       VideoSource = 8006  // Auxilary 6
	VideoSourceME1Prog         VideoSource = 10010 // ME 1 Prog
	VideoSourceME1Prev         VideoSource = 10011 // ME 1 Prev
	VideoSourceME2Prog         VideoSource = 10020 // ME 2 Prog
	VideoSourceME2Prev         VideoSource = 10021 // ME 2 Prev
	VideoSourceCamera1Direct   VideoSource = 11001 // Camera 1 Direct
)
