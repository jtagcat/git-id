package pkg

import (
	"fmt"
	"net/url"

	"github.com/gogs/git-module"
	"github.com/rs/zerolog/log"
	giturls "github.com/whilp/git-urls"
)

func GitRemoteURLGet0(r *git.Repository, remote string) (*url.URL, error) {
	urlString, err := r.RemoteGetURL(remote)
	if err != nil {
		return nil, err
	}
	return giturls.Parse(urlString[0]) // [0]: MVP
}

// if empty, path is working dir
func gitOpen(path string) (*git.Repository, error) {
	// can I haz git?
	_, err := git.BinVersion()
	if err != nil {
		return nil, fmt.Errorf("git binary: %w", err)
	}

	// parse path
	path, err = PWDIfEmpty(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't get working directory: %w", err)
	}

	// open repo
	r, err := git.Open(path)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Debug().Str("path", path).Msg("repo opened")
	return r
}
