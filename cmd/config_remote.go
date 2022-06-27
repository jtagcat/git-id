package cmd

import (
	"fmt"
	"strings"

	valid "github.com/asaskevich/govalidator"
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
				if rs, ok := remoteSlug(fullSlug); ok {
					t.AddRow([]string{
						rs,
						tree.Children.Hostname,
						tree.Children.XDescription,
					})
				}
			}
		}

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
			return fmt.Errorf("expected exactly 2 arguments")
		}

		slug := args.Get(0)

		if strings.HasPrefix(slug, ".") || strings.HasSuffix(slug, ".") {
			return fmt.Errorf("slug can't start or end with a dot")
		}
		if !valid.IsUTFLetterNumeric(
			strings.ReplaceAll(slug, ".", "")) {
			return fmt.Errorf("please choose a saner slug")
		}
		if len(slug) > 128 { // leave some for your userpart asw; not utf-8 len since it cancels out
			return fmt.Errorf("please choose a shorter slug")
		}
		if in := inInvalids(slug); in {
			return fmt.Errorf("slug would conflict with commands, please choose an another (shorter?)")
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
			Hostname:       host,
			IdentitiesOnly: true,
			XDescription:   ctx.String("description"),
		})

		return c.Write()
	},
}

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
			return fmt.Errorf("expected exactly 1 argument")
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
			if t.Children.XParent == suffixSlug || t.Values[0] != "*"+suffixSlug {
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
