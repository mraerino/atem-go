package models

// Topology stores the available number of different features
type Topology struct {
	MEs              uint8
	Sources          uint8
	ColorGenerators  uint8
	AUXBusses        uint8
	DownstreamKeyers uint8 // warning: most likely wrong!
	Stingers         uint8 // warning: most likely wrong!
	DVEs             uint8
	SuperSources     uint8 // warning: most likely wrong!
	SDOutput         bool
}

// MixEffectConfig stores the number of Keyers per ME
type MixEffectConfig map[int]int

// MediaPlayerConfig stores the capacity of the media player buffers
type MediaPlayerConfig struct {
	StillBanks uint8
	ClipBanks  uint8
}
