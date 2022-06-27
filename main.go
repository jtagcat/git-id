package main

import (
	"log"
	"os"

	"github.com/jtagcat/git-id/cmd"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// ~/.ssh/config:
// Include git-id.conf

// git-id.conf:
// # This file is managed by git-id
//
// # This is fully independant of everything else
// Match OriginalHost github.com
//   IdentityFile ~/.ssh/id_rsa
//
// # Remote hooks to actual host
// Host *.gh.git-id
//   Hostname github.com
//   #XDescription "iz GitHub"
//
// # Identity relies on remote
// Host jc.gh.git-id
//  IdentityFile ~/.ssh/id_rsa
//  #XGitConfig user.name jtagcat
//  #XGitConfig user.email blah
//  #XDescription uwu
// Host w.gh.git-id
//  IdentityFile ~/.ssh/work_sk
//
