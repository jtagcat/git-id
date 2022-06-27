package cmd

import "github.com/urfave/cli/v2"

// git-id whoami
var cmdWhoami = &cli.Command{
	Name:      "remote",
	Usage:     "Manage remotes",
	ArgsUsage: "git-id whoami [id] [-1] [-t, --test] [-u, --test-user]",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "1", Aliases: []string{"u"}, Usage: "(without any other options) print only id"},
		&cli.BoolFlag{Name: "test", Aliases: []string{"t"}, Usage: "try to ssh to git@host"},
		&cli.StringFlag{Name: "test-user", Aliases: []string{"u"}, Usage: "try to ssh to <user>@host"},
		flagChdir, flagConfig,
	},
	Hidden: true,
	// -t OR -u OR both: test
	// -u: get non-default user
	// id: use that OR get from pwd / -C
	// ...
}
