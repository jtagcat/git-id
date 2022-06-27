package cmd

import (
	"github.com/urfave/cli/v2"
)

// git-id clone
var cmdClone = &cli.Command{
	Name:      "clone",
	Usage:     "git clone with an identity",
	ArgsUsage: "git-id clone <identity> ",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "1", Aliases: []string{"u"}, Usage: "(without any other options) print only id"},
		&cli.BoolFlag{Name: "test", Aliases: []string{"t"}, Usage: "try to ssh to git@host"},
		&cli.StringFlag{Name: "test-user", Aliases: []string{"u"}, Usage: "try to ssh to <user>@host"},
		flagConfig,
	},
	Hidden: true,
	// get all ids (might be multi-remote)

	// Create temporary ssh-config copying matcidents with special params
	// GIT_SSH_COMMAND="-F <file>" git clone $@
	// Delete file
	// git-id <id> -C <dir>
	// MVP: assume same directory, if isn't, give up with instructions

	// watching dir for before/after would be a good bet (look for new git dir with same origin) (but does not work for absolute paths)
	// parsing all args seems unreasonable, unless we could get git to emit it

	// OR:
	// get all ids where host is same as grep $@
	// use

}

// git remote add origin git@github.com:jtagcat/emptytest.git
