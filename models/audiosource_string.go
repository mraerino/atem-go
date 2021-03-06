// Code generated by "stringer -type=AudioSource -linecomment"; DO NOT EDIT.

package models

import "strconv"

const (
	_AudioSource_name_0 = "Input 1Input 2Input 3Input 4Input 5Input 6Input 7Input 8Input 9Input 10Input 11Input 12Input 13Input 14Input 15Input 16Input 17Input 18Input 19Input 20"
	_AudioSource_name_1 = "XLR"
	_AudioSource_name_2 = "AES/EBU"
	_AudioSource_name_3 = "RCA"
	_AudioSource_name_4 = "Mic 1Mic 2"
	_AudioSource_name_5 = "MP1MP2"
)

var (
	_AudioSource_index_0 = [...]uint8{0, 7, 14, 21, 28, 35, 42, 49, 56, 63, 71, 79, 87, 95, 103, 111, 119, 127, 135, 143, 151}
	_AudioSource_index_4 = [...]uint8{0, 5, 10}
	_AudioSource_index_5 = [...]uint8{0, 3, 6}
)

func (i AudioSource) String() string {
	switch {
	case 1 <= i && i <= 20:
		i -= 1
		return _AudioSource_name_0[_AudioSource_index_0[i]:_AudioSource_index_0[i+1]]
	case i == 1001:
		return _AudioSource_name_1
	case i == 1101:
		return _AudioSource_name_2
	case i == 1201:
		return _AudioSource_name_3
	case 1301 <= i && i <= 1302:
		i -= 1301
		return _AudioSource_name_4[_AudioSource_index_4[i]:_AudioSource_index_4[i+1]]
	case 2001 <= i && i <= 2002:
		i -= 2001
		return _AudioSource_name_5[_AudioSource_index_5[i]:_AudioSource_index_5[i+1]]
	default:
		return "AudioSource(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
