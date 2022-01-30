package ssh_config

import (
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

// general decode/encode: root comment
// root/sub placement change
/*
*  Host
*    Value
*	 XSub
*  XHost // lookahead: if uncommented in same tree
*    XSub
*    InvalidValue // if there are any non-comments between
*  Host

 */

// rootXKeys: which should be considered as TLD (x)keys
// NOTE: rootXKeys may not have any children; to be implemented: x-only children
// unbalanced: trees might not be how they should be, comment placement with newlines
func DecodeXKeysUnbalanced(cfg []RawTopLevel, xkeyPrefix string, rootXKeys []string) (newCfg []RawTopLevel, _ error) {
	rootXKMap := make(map[string]bool)
	for _, rxk := range rootXKeys {
		rootXKMap[strings.ToLower(rxk)] = false
	}
	var iOffset int
rootParsing:
	for i, tree := range cfg {
		newCfg = append(newCfg, RawTopLevel{Key: tree.Key, Values: tree.Values, Comment: tree.Comment, EncodingKVSeperatorIsEquals: tree.EncodingKVSeperatorIsEquals}) // add root without children
		for ci, child := range tree.Children {
			if child.Key == "" && strings.HasPrefix(strings.ToLower(child.Comment), strings.ToLower(xkeyPrefix)) { // xkey
				t, err := decodeLine(child.Comment)
				if err != nil {
					return nil, err
				}

				if _, ok := rootXKMap[strings.ToLower(t.Key)]; ok { // xroot
					// detach
					newCfg = append(newCfg, RawTopLevel{Key: t.Key, Values: t.Values, Comment: t.Comment, EncodingKVSeperatorIsEquals: t.EncodingKVSeperatorIsEquals})
					iOffset++
					// add detached items
					for _, c := range tree.Children[ci+1:] {
						// check that there aren't any non-comments until the next non-xkey,
						// as else we'd move the virtual root of the valid key off its actual root
						if c.Key != "" {
							return nil, ErrValidSubkeyAfterXRoot
						}

						t, err := decodeLine(c.Comment)
						if err != nil {
							return nil, err
						}
						if _, ok := rootXKMap[strings.ToLower(t.Key)]; ok { // xroot
							newCfg = append(newCfg, RawTopLevel{Key: t.Key, Values: t.Values, Comment: t.Comment, EncodingKVSeperatorIsEquals: t.EncodingKVSeperatorIsEquals})
							iOffset++
						} else {
							newCfg[i+iOffset].Children = append(newCfg[i+iOffset].Children, t)
						}
					}
					continue rootParsing // complete detaching
				} // non-root xkey
				newCfg[i+iOffset].Children = append(newCfg[i+iOffset].Children, t)
			} else { // non-keys
				newCfg[i+iOffset].Children = append(newCfg[i+iOffset].Children, child)
			}
		}

		if tree.Key == "" && strings.HasPrefix(strings.ToLower(tree.Comment), strings.ToLower(xkeyPrefix)) {
			t, err := decodeLine(tree.Comment)
			if err != nil {
				return nil, err
			}
			newCfg[i+iOffset] = RawTopLevel{Key: t.Key, Values: t.Values, Comment: t.Comment, EncodingKVSeperatorIsEquals: t.EncodingKVSeperatorIsEquals}
		}
	}
	// TODO: join any root objects with last root, if it was detached by comment parser
	return newCfg, nil
	//TODO: merge this func to DecodeToRaw:
	// 1. if xkey, parse it
	// 2. root xkey logic to root nonxkey logic
	// 3. seperateish non-root xkey logic
	// x. encoder support
}
func balanceDecodedXKeys(cfg []RawTopLevel)

func keywordToTLD(k RawKeyword) RawTopLevel {
	return RawTopLevel{Key: k.Key, Values: k.Values, Comment: k.Comment, EncodingKVSeperatorIsEquals: k.EncodingKVSeperatorIsEquals}
}

// bool: if children were lost
func TLDToKeyword(t RawTopLevel) (RawKeyword, bool) {
	return RawKeyword{Key: t.Key, Values: t.Values, Comment: t.Comment, EncodingKVSeperatorIsEquals: t.EncodingKVSeperatorIsEquals},
		t.Children != nil
}

// func EncodeXKeys(cfg []RawTopLevel, xkeyPrefix string) []RawTopLevel {

// }

//

//

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
// func tldMatch(key string, value []RawValue, search string) (bool, error) {
// 	switch key {
// 	default:
// 		return false, fmt.Errorf("%w: TLD %s", ErrInvalidKeyword, key)
// 	case "Include":
// 		return false, nil
// 	case "Match":
// 		return false, fmt.Errorf("%w: TLD %s", ErrNotImplemented, key)
// 	case "Host":
// 		for _, v := range value {
// 			if strings.EqualFold(v.Value, search) {
// 				return true, nil
// 			}
// 		}
// 		return false, nil
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
