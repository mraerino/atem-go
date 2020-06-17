package atem

import "github.com/mraerino/atem-go/models"

type Characteristics struct {
	ProductName   string
	Topology      models.Topology
	MixEffect     models.MixEffectConfig
	MediaPlayer   models.MediaPlayerConfig
	MultiViews    int
	SuperSources  int // num of boxes
	TallyChannels int
	MacroBanks    int
}

type SwitcherState struct {
	Version struct {
		Major int
		Minor int
	}
	Warning   string
	Power     models.PowerStatus
	VideoMode models.VideoMode

	Inputs map[models.VideoSource]models.InputProperties
	Config Characteristics

	Program map[int]models.VideoSource
	Preview map[int]models.VideoSource
	Aux     map[int]models.VideoSource

	MediaPlayer map[int]*models.MediaPlayer
	MediaFiles  map[int]models.MediaStillFrame
}

func NewSwitcherState() SwitcherState {
	return SwitcherState{
		Inputs: make(map[models.VideoSource]models.InputProperties),
		Config: Characteristics{
			MixEffect: make(models.MixEffectConfig),
		},
		Program: make(map[int]models.VideoSource),
		Preview: make(map[int]models.VideoSource),
		Aux:     make(map[int]models.VideoSource),

		MediaPlayer: make(map[int]*models.MediaPlayer),
		MediaFiles:  make(map[int]models.MediaStillFrame),
	}
}
