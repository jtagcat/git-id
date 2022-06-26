package cmd

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/urfave/cli/v2"
)

// git id remote
var cmdRemote = &cli.Command{
	Name:  "remote",
	Usage: "Manage remotes",
	Subcommands: []*cli.Command{
		cmdRemoteAdd,
	},
}

// git id remote add
var cmdRemoteAdd = &cli.Command{
	Name:      "add",
	Usage:     "Add an remote",
	ArgsUsage: "git-id remote add <remote slug> <actual host> [-d description]",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "description", Aliases: []string{"d"}, Usage: "git-id-only, memory refresher"},
		&cli.PathFlag{Name: "config", Value: "~/.ssh/git-id.conf", Usage: "path to git-id config file"},
	},
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("expected exactly 2 arguments")
		}

		slug := args.Get(0)
		if !valid.IsUTFLetterNumeric(slug) {
			return fmt.Errorf("please choose a saner slug")
		}
		if len(slug) > 200 { // leave some for your userpart asw; not utf-8 len since it cancels out
			return fmt.Errorf("please choose a shorter slug")
		}
		fullSlug := fmt.Sprintf("*.%s.%s", slug, flTLD)

		host := args.Get(1) // don't validate

		//// START ////

		c := gidOpenConfig(ctx.String("config"))

		// Host *.gh.git-id
		if i, trees := c.GID_RootObjectCount("Host", []string{fullSlug}, false); i > 0 {
			return fmt.Errorf("a remote with the slug %s already exists: %s", fullSlug, trees[0].SecondValues)
		}

		c.GIDRootObjectSet("Host", []string{fullSlug, host}, ssh_config.GitIDCommonChildren{
			XDescription: ctx.String("description"),
		})

		return c.Write()
	},
}

var cmdRemoteRemove = &cli.Command{
	Name:      "rm",
	Usage:     "Remove a remote",
	ArgsUsage: "git-id remote rm <remote slug> [--recursive]",
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "recursive", Aliases: []string{"r", "R"}, Usage: "remove remote and associated identities recursively"},
		&cli.PathFlag{Name: "config", Value: "~/.ssh/git-id.conf", Usage: "path to git-id config file"},
	},
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 1 {
			return fmt.Errorf("expected exactly 1 argument")
		}

		slug := args.First()

		c := gidOpenConfig(flConfigPath)

		// Host *.gh.git-id
		fullSlug := fmt.Sprintf("*.%s.%s", slug, flTLD)
		if i, _ := c.GID_RootObjectCount("Host", []string{fullSlug}, false); i == 0 {
			return fmt.Errorf("a remote with the slug %s does not exist", fullSlug)
		}

		// get children
		// Host jc.gh.git-id
		// Match OriginalHost github.com

		// get defaults

		// refuse

		// remove

		// if ok := c.GIDRootObjectRemove()
		return fmt.Errorf("not implemented")
	},
}

// func init() {
// 	cmdRemote.AddCommand(rmRemoteCmd)
// }

// git id remote set: NOVMP
// username
// email

// git id remote rename: VERY NOMVP
