package pkg

import (
	"errors"
	"strings"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
)

var (
	//	ErrTooManyMatches = errors.New("too many matches")
	ErrNoMatch = errors.New("no match")
)

// MVP: will only get the first find
// header is case sensitive
//
// host: jc.gh.git-id; key: "XGitConfig user.name", keyIsComment: true
func KeywordGet(tls []ssh_config.RawTopLevel, header, host, key string, keyIsComment bool) (sliceVals []string, err error) {
	// var hostfinds, keyfinds int
	var rightSection bool
	for _, tl := range tls {
		if !rightSection && header != "" && tl.Comment == header {
			rightSection = true
			continue
		}
		if rightSection || header == "" {
			if strings.EqualFold(tl.Key, "host") { // MVP: Match unsupported
				for _, tv := range tl.Values {
					// TODO: split off to func()?
					if strings.EqualFold(tv.Value, host) {
						// hostfinds++
						// if hostfinds > 1 {
						// 	err = ErrTooManyMatches
						// }
						for _, sv := range tl.Children {
							if !keyIsComment && strings.EqualFold(sv.Key, key) {
								// keyfinds++
								// if keyfinds > 1 {
								// 	err = ErrTooManyMatches
								// }
								for _, svv := range sv.Values {
									sliceVals = append(sliceVals, svv.Value)
								}
								return sliceVals, err
							}
							if keyIsComment && strings.EqualFold(sv.Comment, key) {
								cv, _, err := ssh_config.DecodeValue(sv.Comment)
								if err != nil {
									return sliceVals, err
								}
								for _, cvv := range cv {
									sliceVals = append(sliceVals, cvv.Value)
								}
								return sliceVals, err
							}
						}
					}
				}
			}
		}
	}
	return nil, ErrNoMatch
}

// MVP: substring can't be slice
func FindRemote(tls []ssh_config.RawTopLevel, header, subkey string, substring string, keyIsComment bool) (host []string, err error) {
	// var hostfinds, keyfinds int
	var rightSection bool
	for _, tl := range tls {
		if !rightSection && header != "" && tl.Comment == header {
			rightSection = true
			continue
		}
		if rightSection || header == "" {
			for _, svr := range tl.Children {
				if !keyIsComment && strings.EqualFold(svr.Key, subkey) ||
					keyIsComment && strings.EqualFold(svr.Comment, subkey) {
					if len(svr.Values) != 1 {
						//	return nil, ssh_config.ErrSingleValueOnly
					} else {
						if strings.EqualFold(svr.Values[0].Value, substring) {
							for _, tlvv := range tl.Values {
								host = append(host, tlvv.Value)
							}
							return host, nil
						}
					}

				}
			}
		}
	}
	return nil, ErrNoMatch
}

// FindRemote
// search: HostName github.com; get *.gh-git.id

// KeywordSet
// host: "*.gh.git-id"; key:  "XDescription"; keyIsComment: true

// doing minimum repetitive things: https://www.youtube.com/watch?v=EZ05e7EMOLM

// Host jc.gh.git-id
//  IdentityFile ~/.ssh/id_rsa
//  #XGitConfig user.name jtagcat
//  #XGitConfig user.email blah
//  #XDescription uwu
// Host *.gh.git-id
//   HostName github.com
//   #XDescription "iz GitHub"
