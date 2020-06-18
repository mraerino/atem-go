package models

type PowerStatus struct {
	Main   bool
	Backup bool
}

type TallyState struct {
	Program bool
	Preview bool
}

// Bitmask is the on-wire representation of this state
func (t TallyState) Bitmask() uint8 {
	var mask uint8
	if t.Program {
		mask |= 1
	}
	if t.Preview {
		mask |= 2
	}
	return mask
}

type Timecode struct {
	Hour   int
	Minute int
	Second int
	Frame  int
}
