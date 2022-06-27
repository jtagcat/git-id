package cmd

import (
	"github.com/urfave/cli/v2"
)

// git-id config id
var cmdConfigId = &cli.Command{
	Name:    "id",
	Aliases: []string{"identity"},
	Subcommands: []*cli.Command{
		cmdConfigIdAdd,
		cmdConfigIdRm,
	},
}

// git-id add
var cmdConfigIdAdd = &cli.Command{
	Name:      "add",
	Usage:     "Add an identity",
	ArgsUsage: "git-id config id add <remote> <identity> <IdentityFile> [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
	Flags: []cli.Flag{
		flagDesc,
		&cli.StringFlag{Name: "username", Aliases: []string{"u"}, Usage: "git config user.name"},
		&cli.StringFlag{Name: "email", Aliases: []string{"e"}, Usage: "git config user.email"},
		&cli.StringFlag{Name: "signing-key", Aliases: []string{"sk"}, Usage: "git config user.signingKey"},
		flagConfig,
	},
	Hidden: true,
	// -t OR -u OR both: test
	// -u: get non-default user
	// id: use that OR get from pwd / -C
	// ...
}
