package models

//go:generate go run golang.org/x/tools/cmd/stringer -type=MediaPlayerType -linecomment -output=media_string.go
type MediaPlayerType uint8

const (
	MediaPlayerTypeStill MediaPlayerType = 1 + iota // Still
	MediaPlayerTypeClip                             // Clip
)

type MediaPlayer struct {
	Type       MediaPlayerType
	StillIndex int
	ClipIndex  int
	// missing: clip player state (RCPS)
}

type MediaStillFrame struct {
	Used     bool
	Hash     []byte
	Filename string
}
