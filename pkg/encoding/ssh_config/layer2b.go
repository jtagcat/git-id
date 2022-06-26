package ssh_config

import (
	"strings"

	"github.com/minio/pkg/wildcard"
)

// Splits Config to (up to) 3, based on start/end headers (usually xkeys).
//
// Match is considered when any of the header key-value is equal case-insensitivly.
//
// Both inputs are optional.
func (c *Config) getBetweenHeaders(startHeaders, eofHeaders []RawTopLevel) (beforeHeader, betweenHeaders, afterHeader []RawTopLevel) {
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
	for _, line := range c.cfg {
		switch state {
		case 0: // before header
			if values, ok := startHeaderMap[strings.ToLower(line.Key)]; ok { // key matches
				if rawValuesMatch(line.Values, values) {
					// found a starting header
					state = 1
					continue
				}
			}
			beforeHeader = append(beforeHeader, line)
		case 1: // during header
			if values, ok := eofHeaderMap[strings.ToLower(line.Key)]; ok {
				if rawValuesMatch(line.Values, values) {
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

func rawValuesMatch(a, b []RawValue) bool { // similar to EqualFoldSlice
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

func oneRawValuesMatch(a []RawValue, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !strings.EqualFold(v.Value, b[i]) {
			return false
		}
	}
	return true
}

// get all roots what have defined KV under them
func RootBySubKV(cfg []RawTopLevel, rootKey, subKey string, subValue []string) (scfg []RawTopLevel) {
	for _, r := range cfg {
		if strings.EqualFold(r.Key, rootKey) {
			for _, sk := range r.Children {
				if strings.EqualFold(sk.Key, subKey) {
					if oneRawValuesMatch(sk.Values, subValue) {
						scfg = append(scfg, r)
					}
				}
			}
		}
	}
	return scfg
}

// []RawTopLevel

func rawValuesToStrSlice(raw []RawValue) (values []string) {
	for _, w := range raw {
		values = append(values, w.Value)
	}
	return values
}

// if any ! matches, return false
// jfyi Host hello,world foo is "hello,world", "foo"
func HostMatches(raw []string, search string) (matches bool) {
	locaseSearch := strings.ToLower(search)
	for _, w := range raw {
		if strings.HasPrefix(w, "!") { // negating pattern
			if wildcard.Match(strings.ToLower(w[1:]), locaseSearch) {
				// negating pattern matches
				return false // exit early
			}
		} else if !matches {
			matches = wildcard.Match(strings.ToLower(w), locaseSearch)
			// if match found, continue searching for negating patterns
		}
	}
	return matches
}

var matchKeywords = []string{"canonical", "exec", "host", "originalhost", "user", "localuser"}

// specials: all, final

// err: nil, ErrInvalidValue
// func buildMatchTree(raw []string) (tree []RawKeyword, err error) {
// 	var keywordActive bool
// 	for _, w := range raw {
// 		// in "Match=key=value,bar don" / "Match key=value,bar don"
// 		// raw is {"key=value,bar", "don"}, w is "key=value,bar", d is Key: "key", []values:{"value", "bar"}
// 		// quotes and escapes are already handled by lower level decode/encode
// 		d, _ := decodeLine(w)
// 		if d.Comment != "" {
// 			// when parent of nested decode is quoted
// 			d.Values = append(d.Values, RawValue{d.Comment, 0})
// 		}

// 		if !d.EncodingKVSeperatorIsEquals {
// 			if d.Values != nil {
// 				// when user gives "key=value,bar don" or
// 				// "key value,bar", "don" or "key value bar", "don" (all invalid)
// 				// instead of "key=value,bar", "don"
// 				return nil, fmt.Errorf("value must not have unquoted spaces in them, "+
// 					"it should have been decoded by lower-stage decoder "+
// 					"or use '=' as a subkey-subvalue denominator: %w", ErrInvalidValue)
// 			}
// 			// valueless keyword
// 			//TODO:
// 			//
// 			keywordActive = true
// 			continue
// 		}
// 		// valueful keyword

// 	}
// 	return tree, nil
// }

// func MatchSubValueMatches(raw []string, subkey, search string) bool {
// 	for i, w := range raw {
// 		k, _ := decodeLine(w)
// 		if k.Comment != "" {
// 			continue // invalid, can't have '#' in a single value, it should already be parsed
// 		}
// 		if strings.EqualFold(k.Key, subkey) {
// 			if k.EncodingKVSeperatorIsEquals {
// 				// k.Values is csvStringSlice
// 			} else {

// 			}
// 		}
// 	}
// }

func MatchMatches(raw []string, search string) bool {
	return false
}
