package cmd

import (
	"github.com/urfave/cli/v2"
)

// git-id add
var cmdSet = &cli.Command{
	Name:      "set",
	Usage:     "Add an identity",
	ArgsUsage: "git-id config id <remote> <identity> [-i IdentityFile] [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
	Flags: []cli.Flag{
		&cli.PathFlag{Name: "username", Aliases: []string{"u"}, Usage: "git config user.name"},
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

	// get is basic whoami
}
