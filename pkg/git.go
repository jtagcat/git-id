package pkg

import (
	"net/url"

	"github.com/gogs/git-module"
	giturls "github.com/whilp/git-urls"
)

func GitRemoteURLGet0(r *git.Repository, remote string) (*url.URL, error) {
	urlString, err := r.RemoteGetURL(remote)
	if err != nil {
		return nil, err
	}
	return giturls.Parse(urlString[0]) // [0]: MVP
}
