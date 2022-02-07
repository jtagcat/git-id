package cmd

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/rs/zerolog/log"
)

// nil, fs.ErrInvalid
func giRemotesFromAnyHost(cfg []ssh_config.RawTopLevel, host string) (remotes []string, err error) {
	if host == "" {
		return nil, fmt.Errorf("%w: host is empty", fs.ErrInvalid)
	}

	hostParts := strings.Split(host, ".")
	tld := hostParts[len(hostParts)-1]
	if tld == flGI_TLD { // iz already git-id
		if len(hostParts) != 3 {
			return nil, fmt.Errorf("%w: host has git-id TLD, but is does not have 3 parts", fs.ErrInvalid)
		}
		// we have user.remote.git-id, forward remote
		//TODO: potential bugish behaviour: from x.a.git-id to y.b.git-id, where b and a have the same underlieing domain
		// var a: intended to see them as different (preferred)
		// var b: get remote Hostname, and do same as non-hijcacked hosts
		return []string{hostParts[1]}, nil
	}

	// non-hijacked host
	gitidRemotes := ssh_config.RootBySubKV(cfg, "Host", "Hostname", []string{host})
	if len(gitidRemotes) == 0 {
		log.Fatal().Str("host", host).Msg("no git-id remote for host found")
	}
	// we got a list of "*.xyz.git-id"
	// extract git-id remote and forward
	for _, idr := range gitidRemotes {
		for _, v := range idr.Values {
			giHostParts := strings.Split(v.Value, ".")
			if len(giHostParts) == 3 && giHostParts[0] == "*" && giHostParts[2] == flGI_TLD {
				remotes = append(remotes, giHostParts[1])
			}
		}
	}
	return remotes, nil
}
