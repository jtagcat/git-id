package cmd

import (
	"log"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/urfave/cli/v2"
)

var (
	gitidHeaderInfo         = " This file is managed by git-id"
	gitidHeaderDefaults     = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Default identities", Quoted: 0}}}
	gitidHeaderIdentities   = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Identities", Quoted: 0}}}
	gitidHeaderRemotes      = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Remotes", Quoted: 0}}}
	gitidSSHConfigRootXKeys = map[string]bool{"xheader": false}
	gitidSSHConfigSubXKeys  = []string{"XGitConfig", "XDescription"}

	userSSHConfigFile = "~/.ssh/config"
	remote            = "origin"
)

func Execute(args []string) {
	if err := app.Run(args); err != nil {
		log.Fatalln(err)
	}
}

var app = &cli.App{
	Name:  "git-id",
	Usage: "Git identity management",
	Flags: []cli.Flag{
		&cli.PathFlag{Name: "config", Value: "~/.ssh/git-id.conf", Usage: "path to git-id config file"},
	},
	Commands: []*cli.Command{
		cmdRemote,
	},
}

// NOTMVP: git branch, ncdu-style, whatever arrow keys / fzf / quick switcher

var (
	flTLD        = "git-id"
	flConfigPath string
	flActPath    string
)

// func init() {
// 	pkg.ZerologLevelStringint(os.Getenv("LOGLEVEL")) // TODO: parse -vvv and --verbose=5 / --verbose=info
// 	rootCmd.PersistentFlags().StringVarP(&flActPath, "", "C", "", "Act on `path` instead of working directory.") //**HACK1** bugbug upstream: https://github.com/spf13/pflag/issues/139
// }

// func accessConfig() *os.File {
// 	//
// }

// NOTMVP: custom core.sshCommand additions
// very NOMVP: allow hiding/deprecating an id/remote instead of rm
// TODO: rm/change/deprecate: can we use ssh_config things to print something / execute git-id hidden command?

// gitidConfig:
// # This file is managed by git-id
//
// #XHeader Defaults
// Match OriginalHost github.com
//   IdentityFile ~/.ssh/id_rsa
//
// #XHeader Remotes
// Host *.gh.git-id
//   Hostname github.com
//   #XDescription "iz GitHub"
//
// #XHeader Identities
// Host jc.gh.git-id
//  IdentityFile ~/.ssh/id_rsa # this is redundant with defaults, IdentityFile is used for matching the default to an ident
//  #XGitConfig user.name jtagcat
//  #XGitConfig user.email blah
//  #XDescription uwu
// Host w.gh.git-id
//  IdentityFile ~/.ssh/work_sk
//
