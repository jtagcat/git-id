package cmd

import (
	"fmt"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/urfave/cli/v2"
)

// upstream bug: https://github.com/urfave/cli/issues/1217
var cmdDefaultUsageBug = "git-id default <host> <default IdentityFile>"

// git-id default
var cmdDefault = &cli.Command{
	Name:  "default",
	Usage: "Set host's default key",
	Subcommands: []*cli.Command{
		cmdDefaultClear,
	},

	ArgsUsage: cmdDefaultUsageBug,
	Flags: []cli.Flag{
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 2 {
			// upstream bug: https://github.com/urfave/cli/issues/1217
			// fmt.Println("Usage:", ctx.Command.ArgsUsage)
			fmt.Println(cmdDefaultUsageBug)
			return fmt.Errorf("expected exactly 2 arguments")
		}

		host := args.Get(0)
		idfile := args.Get(1)

		c := gidOpenConfig(ctx.String("config"))

		// Match OriginalHost github.com
		c.GID_RootObjectSetFirst("Match", []string{"OriginalHost", host}, true, ssh_config.GitIDCommonChildren{
			IdentityFile:   idfile,
			IdentitiesOnly: true,
		})

		return c.Write()
	},
}

// git-id default clear
var cmdDefaultClear = &cli.Command{
	Name:      "clear",
	Usage:     "Clear previously set default key",
	ArgsUsage: "git-id default clear <host>",
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 1 {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("expected exactly 1 argument")
		}

		host := args.Get(0)

		c := gidOpenConfig(ctx.String("config"))

		// Match OriginalHost github.com
		if ok := c.GIDRootObjectRemoveFirst("Match", []string{"OriginalHost", host}); !ok {
			return fmt.Errorf("host %s does not have a default key set", host)
		}

		return c.Write()
	},
}
