package ssh_config

import (
	"strings"
)

// splits []RawTopLevel to up to 3, based on start/end header key-values (headers are usually comments, use xkeys).
// 2nd return is what you usually want.
//
// Only the first match will be considered;
// header of default value ("") will be ignored.
func GetBetweenHeaders(cfg []RawTopLevel, startHeaders, eofHeaders []RawTopLevel) (beforeHeader, betweenHeaders, afterHeader []RawTopLevel) {
	startHeaderMap := make(map[string][]RawValue)
	for _, h := range startHeaders {
		startHeaderMap[strings.ToLower(h.Key)] = h.Values
	}
	eofHeaderMap := make(map[string][]RawValue)
	for _, h := range eofHeaders {
		eofHeaderMap[strings.ToLower(h.Key)] = h.Values
	}

	var state int
	if len(startHeaders) == 0 {
		state = 1
	}
	for _, line := range cfg {
		switch state {
		case 0: // before header
			if values, ok := startHeaderMap[strings.ToLower(line.Key)]; ok { // key matches
				if valuesMatch(line.Values, values) {
					// found a starting header
					state = 1
					continue
				}
			}
			beforeHeader = append(beforeHeader, line)
		case 1: // during header
			if values, ok := eofHeaderMap[strings.ToLower(line.Key)]; ok {
				if valuesMatch(line.Values, values) {
					// found an eofHeader
					state = 2
					continue
				}
			}
			betweenHeaders = append(betweenHeaders, line)
		case 2: // after eofHeader
			afterHeader = append(afterHeader, line)
		}
	}
	return beforeHeader, betweenHeaders, afterHeader
}

func valuesMatch(a, b []RawValue) bool { // similar to EqualFoldSlice
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !strings.EqualFold(v.Value, b[i].Value) {
			return false
		}
	}
	return true
}

//

//

// func TLDofKV(cfg []RawTopLevel, subkeyIsComment bool, subkey string, subvalues []string) {
// 	for _, line := range cfg {
// 		for _, kv := range line.Children {
// 			if !subkeyIsComment && strings.EqualFold(kv.)
// 		}
// 	}
// }

// request: specific key(s) under a specified host(s) jq: '.[]host | | select(.key == $key)'
// returns slice of matches (rawKeyword)
// not intended: using the returned values directly for writing (use nth)
// MATCH not implemeneted
// func MatchingHostsTrees(cfg []RawTopLevel, hostkey, hostvalue, key string, keyIsComment bool) {
// }
