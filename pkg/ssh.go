package pkg

//
// import (
// 	"errors"
// 	"strings"
//
// 	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
// )
//
// var (
// 	//	ErrTooManyMatches = errors.New("too many matches")
// 	ErrNoMatch = errors.New("no match")
// )
//
// // MVP: will only get the first find
// // header is case sensitive
// //
// // host: jc.gh.git-id; key: "XGitConfig user.name", keyIsComment: true
// //TODO: return []rawKeyword
// //TODO: error handling?
// func HostKeyword(tls []ssh_config.RawTopLevel, header, host, key string, keyIsComment bool) (sliceVals []string, err error) {
// 	// var hostfinds, keyfinds int
// 	var rightSection bool
// 	for _, tl := range tls {
// 		if header != "" && !rightSection && tl.Comment == header {
// 			rightSection = true
// 			continue
// 		}
// 		if header == "" || rightSection {
// 			if strings.EqualFold(tl.Key, "host") { // MVP: Match unsupported
// 				for _, tv := range tl.Values {
// 					// TODO: split off to func()?
// 					if strings.EqualFold(tv.Value, host) {
// 						// hostfinds++
// 						// if hostfinds > 1 {
// 						// 	err = ErrTooManyMatches
// 						// }
// 						for _, sv := range tl.Children {
// 							if !keyIsComment && strings.EqualFold(sv.Key, key) {
// 								// keyfinds++
// 								// if keyfinds > 1 {
// 								// 	err = ErrTooManyMatches
// 								// }
// 								for _, svv := range sv.Values {
// 									sliceVals = append(sliceVals, svv.Value)
// 								}
// 								return sliceVals, err
// 							}
// 							if keyIsComment && strings.EqualFold(sv.Comment, key) {
// 								cv, _, err := ssh_config.DecodeValue(sv.Comment)
// 								if err != nil {
// 									return sliceVals, err
// 								}
// 								for _, cvv := range cv {
// 									sliceVals = append(sliceVals, cvv.Value)
// 								}
// 								return sliceVals, err
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil, ErrNoMatch
// }
//
// // SetRawKV
// // host: "*.gh.git-id"; key:  "XDescription"; keyIsComment: true
// // host: "" for TLD
// func SetRawKV(tls []ssh_config.RawTopLevel, header string, hostV []string, newKV ssh_config.RawKeyword, nth int, tlClearChildren bool) ([]ssh_config.RawTopLevel, error) {
// 	rightSection, currentNth, success := false, 0, false
// setRoutine:
// 	for _, tl := range tls {
// 		if header != "" && !rightSection && tl.Comment == header {
// 			rightSection = true
// 			continue
// 		}
// 		if header == "" || rightSection {
// 			if hostV != nil && strings.EqualFold(tl.Key, newKV.Key) {
// 				if currentNth == nth {
// 					tl.Key = newKV.Key
// 					tl.Values = newKV.Values
// 					tl.Comment = newKV.Comment
// 					tl.EncodingKVSeperatorIsEquals = newKV.EncodingKVSeperatorIsEquals
// 					if tlClearChildren {
// 						tl.Children = nil
// 					}
// 					success = true
// 					break setRoutine
// 				}
// 				currentNth++
// 			} else if hostV == nil {
// 				tlv := getRawValueValues(tl.Values)
// 				if EqualFoldSlice(tlv, hostV) {
// 					for _, tc := range tl.Children {
// 						if strings.EqualFold(tc.Key, newKV.Key) {
// 							if currentNth == nth {
// 								tc = newKV
// 								success = true
// 								break setRoutine
// 							}
// 							currentNth++
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	var err error
// 	if !success {
// 		err = ErrNoMatch
// 	}
// 	return tls, err
// }
//
// func getRawValueValues(rv []ssh_config.RawValue) (values []string) {
// 	for _, rvv := range rv {
// 		values = append(values, rvv.Value)
// 	}
// 	return values
// }
//
// // doing minimum repetitive things: https://www.youtube.com/watch?v=EZ05e7EMOLM
//
// //
// //
// //
//
// // get gitid remotes: from section gitidHeaderRemotes + RootBySubKV
// // get gitid idents of remote(s): from section gitidHeaderIdentities +
