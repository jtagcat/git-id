package ssh_config

import (
	"strings"
)

// splits []RawTopLevel to up to 3, based on start/end header (comment). 2nd is what you usually want.
// Only the first match will be considered;
// header of default value ("") will be ignored.
func GetBetweenHeaders(cfg []RawTopLevel, headers, eofHeaders []string) (beforeHeader, betweenHeaders, afterHeader []RawTopLevel) {
	headerMap := make(map[string]bool)
	for _, h := range headers {
		headerMap[strings.ToLower(h)] = false
	}
	eofHeaderMap := make(map[string]bool)
	for _, h := range eofHeaders {
		eofHeaderMap[strings.ToLower(h)] = false
	}

	var state int
	if len(headers) == 0 {
		state = 1
	}
	for _, line := range cfg {
		switch state {
		case 0: // before header
			if _, ok := headerMap[strings.ToLower(line.Comment)]; ok {
				// found first header
				state = 1
				continue
			}
			beforeHeader = append(beforeHeader, line)
		case 1: // during header
			if _, ok := eofHeaderMap[strings.ToLower(line.Comment)]; ok {
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
// func MatchingHosts(cfg []RawTopLevel, hostkey, hostvalue, key string, keyIsComment bool) {
// 	for _, line := range cfg {
// 		if strings.EqualFold(line.Key) {
// 	}
// }
