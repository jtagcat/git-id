package cmd

import (
	"fmt"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/urfave/cli/v2"
)

// git-id set-default
var cmdSetDefault = &cli.Command{
	Name:      "set-default",
	Usage:     "Set default ssh key for host (not remote), or clear it",
	ArgsUsage: "git-id set-default <host> <IdentityFile>",
	Flags: []cli.Flag{
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 2 {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("expected exactly 2 arguments (hint: to clear, use \"\"")
		}

		host := args.Get(0)
		idfile := args.Get(1)

		c := gidOpenConfig(ctx.Path("config"))

		if idfile == "" { // clear
			// Match OriginalHost github.com
			if ok := c.GID_RootObjectRemoveFirst("Match", []string{"OriginalHost", host}); !ok {
				return fmt.Errorf("host %s does not have a default key set", host)
			}
			return c.Write()
		}

		// Set
		// Match OriginalHost github.com
		c.GID_RootObjectSetFirst("Match", []string{"OriginalHost", host}, true, ssh_config.GitIDCommonChildren{
			IdentityFile:   idfile,
			IdentitiesOnly: true,
		})

		return c.Write()
	},
}
