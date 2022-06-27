package cmd

import (
	"fmt"
	"strings"

	"github.com/gravitational/teleport/lib/asciitable"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/urfave/cli/v2"
)

// git-id remote
var cmdRemote = &cli.Command{
	Name:  "remote",
	Usage: "Manage remotes",
	Subcommands: []*cli.Command{
		cmdRemoteAdd,
		cmdRemoteRemove,
	},
	Flags: []cli.Flag{
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		if ctx.Args().Len() != 0 {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("expected no arguments")
		}

		c := gidOpenConfig(ctx.Path("config"))

		_, trees := c.GID_RootObjects("Host", []string{"." + globalTLD}, true)

		t := asciitable.MakeTable([]string{"Remote", "Host", "Description"})
		for _, tree := range trees {
			fullSlug := tree.Values[0]

			if strings.HasPrefix(fullSlug, "*.") {
				t.AddRow([]string{
					remoteSlug(fullSlug),
					tree.Children.Hostname,
					tree.Children.XDescription,
				})
			}
		}

		t.DiscardEmpty()
		fmt.Println(t.AsBuffer().String())
		return nil
	},
}

// git-id remote add
var cmdRemoteAdd = &cli.Command{
	Name:      "add",
	Usage:     "Add a remote",
	ArgsUsage: "git-id remote add <remote slug> <actual host> [-d description]",
	Flags: []cli.Flag{
		flagDesc,
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 2 {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("expected exactly 2 arguments, got %d", args.Len())
		}

		slug := args.Get(0)
		if err := validateSlug(slug); err != nil {
			return err
		}

		fullSlug := fmt.Sprintf("*.%s.%s", slug, globalTLD)

		host := args.Get(1) // don't validate

		//// START ////

		c := gidOpenConfig(ctx.Path("config"))

		// Host *.gh.git-id
		if i, trees := c.GID_RootObjects("Host", []string{fullSlug}, false); i > 0 {
			return fmt.Errorf("a remote with the slug %s already exists: %s", fullSlug, trees[0].Values)
		}

		c.GID_RootObjectSetFirst("Host", []string{fullSlug}, false, ssh_config.GitIDCommonChildren{
			Hostname:     host,
			XDescription: ctx.String("description"),
		})

		return c.Write()
	},
}

// NONMVP: it'd be nice if it showed u a list of ids + repos :)
// git-id remote rm
var cmdRemoteRemove = &cli.Command{
	Name:      "rm",
	Usage:     "Remove a remote",
	ArgsUsage: "git-id remote rm <remote slug> <-y, --yes> [-r, --recursive]",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "yes", Aliases: []string{"-y"}, Usage: "acknowledge potential breakage"},
		&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}, Usage: "remove remote and associated identities recursively"},
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		if !ctx.Bool("yes") {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("After a remote is removed, all git repos using it will hold broken links. (-y, --yes)")
		}

		//// ARGS ////
		args := ctx.Args()
		if args.Len() != 1 {
			fmt.Println("Usage:", ctx.Command.ArgsUsage)
			return fmt.Errorf("expected exactly 1 argument, got %d", args.Len())
		}

		recursive := ctx.Bool("recursive")

		slug := args.First()
		suffixSlug := fmt.Sprintf("%s.%s", slug, globalTLD)

		c := gidOpenConfig(ctx.Path("config"))

		// Host *.gh.git-id
		i, trees := c.GID_RootObjects("Host", []string{"." + suffixSlug}, true)
		if i == 0 {
			return fmt.Errorf("remote %s does not exist", slug)
		}

		// get/remove children identities
		// Host jc.gh.git-id
		for _, t := range trees {
			if t.Children.XParent == suffixSlug ||
				t.Values[0] != "*"+suffixSlug { // for legacy / just in case there is no parent

				if !recursive {
					return fmt.Errorf("cannot delete remote %s: has attached identities (use --recursive)", slug)
				}

				if ok := c.GID_RootObjectRemoveFirst("Host", t.Values); !ok {
					return fmt.Errorf("race‽ (report bug?): identity doesn't exist, but it just did")
				}
			}
		}

		// remove remote
		if ok := c.GID_RootObjectRemoveFirst("Host", []string{"*" + suffixSlug}); !ok {
			return fmt.Errorf("race‽ (report bug?): remote doesn't exist, but it just did")
		}

		return c.Write()
	},
}
