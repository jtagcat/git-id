package cmd

import "github.com/urfave/cli/v2"

// git-id whoami
var cmdWhoami = &cli.Command{
	Name:      "remote",
	Aliases:   []string{"who"},
	Usage:     "Manage remotes",
	ArgsUsage: "git-id whoami [id] [-1] [-t, --test]",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "1", Aliases: []string{"u"}, Usage: "(without any other options) print only id"},
		&cli.BoolFlag{Name: "test", Aliases: []string{"t"}, Usage: "try to ssh to git@host"},
		flagChdir, flagConfig,
	},
	Hidden: true,
	// -t OR -u OR both: test
	// -u: get non-default user
	// id: use that OR get from pwd / -C
	// ...
}
