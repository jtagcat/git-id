package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	gitidHeaderInfo = " This file is managed by git-id"

	// implementing headers would probably be best with virtual configs
	// gitidHeaderDefaults     = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Default identities", Quoted: 0}}}
	// gitidHeaderIdentities   = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Identities", Quoted: 0}}}
	// gitidHeaderRemotes      = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{Value: "Remotes", Quoted: 0}}}
	// gitidSSHConfigRootXKeys = map[string]bool{"xheader": false}

	// hardcodes
	globalTLD         = "git-id"
	userSSHConfigFile = "~/.ssh/config"

	remote = "origin"
)

// TODO: mention [-C] [-c, --config], putting it everywhere clutters
var App = &cli.App{
	Name:      "git-id",
	Usage:     "Stupid git identity management (list | switch <id>)",
	ArgsUsage: "git-id <identity>", // | [-a, --all]",
	Flags: []cli.Flag{
		// &cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "list all identities (instead of current remote)"},
		flagConfig,
	},
	Commands: []*cli.Command{
		// action
		cmdClone,
		// switch is root
		cmdWhoami,

		// config
		cmdAdd,
		cmdSet,
		cmdRemove,

		cmdRemote,
		cmdSetDefault,

		// help
		cmdHow,
	},
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		switch args.Len() {
		default:
			return fmt.Errorf("expected 0 or 1 arguments")
		case 0:
			return cmdRoot(ctx)
		case 1:
			return cmdSwitch(ctx, args.First())
		}
	},
}

var (
	flagDesc = &cli.StringFlag{Name: "description", Aliases: []string{"d"}, Usage: "git-id-only, memory refresher"}

	flagChdir  = &cli.PathFlag{Name: "C", Usage: "act on git repo at specified path"}
	flagConfig = &cli.PathFlag{Name: "config", Value: "~/.ssh/git-id.conf", Usage: "path to git-id config file"}
)
