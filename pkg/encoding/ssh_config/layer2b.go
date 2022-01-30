package ssh_config

import (
	"fmt"
	"strings"
)

// splits []RawTopLevel to up to 3, based on start/end header (comment). 2nd is what you usually want.
// Only the first match will be considered;
// header of default value ("") will be ignored.
func GetBetweenHeaders(cfg []RawTopLevel, header, eofHeader string) (beforeHeader, betweenHeaders, afterHeader []RawTopLevel) {
	var state int
	if header == "" {
		state = 1
	}
	for _, line := range cfg {
		switch state {
		case 0: // before header
			if line.Key == "" && strings.EqualFold(line.Comment, header) {
				// found first header
				state = 1
				continue
			}
			beforeHeader = append(beforeHeader, line)
		case 1: // during header
			if line.Key == "" && strings.EqualFold(line.Comment, eofHeader) && eofHeader != "" {
				state = 2
				continue
			}
			betweenHeaders = append(betweenHeaders, line)
		case 2: // after eofHeader
			afterHeader = append(afterHeader, line)
		}
	}
	return beforeHeader, betweenHeaders, afterHeader
}

// func TLDofKV(cfg []RawTopLevel, subkeyIsComment bool, subkey string, subvalues []string) {
// 	for _, line := range cfg {
// 		for _, kv := range line.Children {
// 			if !subkeyIsComment && strings.EqualFold(kv.)
// 		}
// 	}
// }

// key enum: Host, Match, Include
// key-value likely part of RawTopLevel
// search is 2nd side of equals
func tldMatch(key string, value []RawValue, search string) (bool, error) {
	switch key {
	default:
		return false, fmt.Errorf("%w: TLD %s", ErrInvalidKeyword, key)
	case "Include":
		return false, nil
	case "Match":
		return false, fmt.Errorf("%w: TLD %s", ErrNotImplemented, key)
	case "Host":
		for _, v := range value {
			if strings.EqualFold(v.Value, search) {
				return true, nil
			}
		}
		return false, nil
	}
}

// request: specific key(s) under a specified host(s) jq: '.[]host | | select(.key == $key)'
// returns slice of matches (rawKeyword)
// not intended: using the returned values directly for writing (use nth)
// func MatchingHosts(cfg []RawTopLevel, hostkey, hostvalue, key string, keyIsComment bool) {
// 	for _, line := range cfg {
// 		if strings.EqualFold(line.Key) {
// 	}
// }
