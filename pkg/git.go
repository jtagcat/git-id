package pkg

import (
	"fmt"
	"net/url"

	"github.com/gogs/git-module"
	"github.com/mitchellh/go-homedir"
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
func gitOpen(name string) (*git.Repository, error) {
	path, err := homedir.Expand(name)
	if err != nil {
		return nil, fmt.Errorf("expanding path %s: %w", name, err)
	}

	// can I haz git?
	if _, err := git.BinVersion(); err != nil {
		return nil, fmt.Errorf("git binary: %w", err)
	}

	// parse path
	path, err = PWDIfEmpty(path)
	if err != nil {
		return nil, fmt.Errorf("getting working directory: %w", err)
	}

	// open repo
	r, err := git.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening git repo: %w", err)
	}
	return r, nil
}
