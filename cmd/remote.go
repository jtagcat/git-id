package cmd

import (
	"fmt"

	valid "github.com/asaskevich/govalidator"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/spf13/cobra"
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

		if i, secondValues := c.GID_RootObjectCount("Host", []string{fullSlug}); i > 0 {
			return fmt.Errorf("a remote with the slug %s already exists: %s", fullSlug, secondValues)
		}

		c.GIDRootObjectSet("Host", []string{fullSlug, host}, ssh_config.GitIDCommonChildren{
			XDescription: ctx.String("description"),
		})

		return c.Write()
	},
}

var flDescription string

// func init() {
// 	addRemoteCmd.LocalFlags().StringVarP(&flDescription, "description", "d", "", "git-id-only, memory refresher")
// }

// git id remote rm: NOMVP
var rmRemoteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an remote",
	Long: `Usage: git-id remote rm <remote slug>
	Example: git-id remote rm gh`,
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected exactly 1 argument")
		}
		fullSlug := fmt.Sprintf("*.%s.%s", args[0], flTLD)

		c := gidOpenConfig(flConfigPath)

		if i, _ := c.GID_RootObjectCount("Host", []string{fullSlug}); i == 0 {
			return fmt.Errorf("a remote with the slug %s does not exist", fullSlug)
		}

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
