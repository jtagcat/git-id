package cmd

import (
	"fmt"
	"strings"

	"github.com/gravitational/teleport/lib/asciitable"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	spkg "github.com/jtagcat/git-id/pkg/encoding/ssh_config/pkg"
	"github.com/urfave/cli/v2"
)

// git-id
func cmdRoot(ctx *cli.Context) error {
	c := gidOpenConfig(ctx.Path("config"))

	_, trees := c.GID_RootObjects("Host", []string{"." + globalTLD}, true)

	hostMap := make(map[string]string)
	for _, tree := range trees {
		fullSlug := tree.Values[0]
		if strings.HasPrefix(fullSlug, "*.") {
			if rs := remoteSlug(fullSlug); rs != "" {
				hostMap[rs] = tree.Children.Hostname
			}
		}
	}

	t := asciitable.MakeTable([]string{"Identity", "Remote", "Real host", "IdentityFile", "g: user.name", "g: user.email", "g: user.signingkey", "Description"})

	// if !ctx.Bool("all") {
	// 	// NONMVP: get git repo, origin, host, remotes matching host, ids matching host
	// }

	for _, tree := range trees {
		fullSlug := tree.Values[0]
		if !strings.HasPrefix(fullSlug, "*.") {
			slug := strings.TrimSuffix(fullSlug, "."+tree.Children.XParent)
			rslug, _, _ := spkg.CutLast(tree.Children.XParent, ".")
			host, ok := hostMap[rslug]
			if !ok {
				host = "<not found>"
			}

			t.AddRow([]string{
				slug, rslug, host, tree.Children.IdentityFile,
				tree.Children.XGitConfigUsername, tree.Children.XGitConfigUserEmail, tree.Children.XGitConfigSigningKey,
				tree.Children.XDescription,
			})
		}
	}

	t.DiscardEmpty()
	fmt.Println(t.AsBuffer().String())
	return nil
}

var (
	flagGitUsername   = &cli.StringFlag{Name: "username", Aliases: []string{"u"}, Usage: "git config user.name"}
	flagGitEmail      = &cli.StringFlag{Name: "email", Aliases: []string{"e"}, Usage: "git config user.email"}
	flagGitSigningkey = &cli.StringFlag{Name: "signing-key", Aliases: []string{"sk"}, Usage: "git config user.signingKey"}
)

// git-id add
var cmdAdd = &cli.Command{
	Name:      "add",
	Usage:     "Add an identity",
	ArgsUsage: "git-id add <remote> <identity> <IdentityFile> [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
	Flags: []cli.Flag{
		flagDesc,
		flagGitUsername, flagGitEmail, flagGitSigningkey,
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 3 {
			return fmt.Errorf("expected exactly 3 arguments, got %d", args.Len())
		}

		c := gidOpenConfig(ctx.Path("config"))

		rslug := args.Get(0)
		rSuffixSlug := fmt.Sprintf("%s.%s", rslug, globalTLD)
		if i, _ := c.GID_RootObjects("Host", []string{"*." + rSuffixSlug}, false); i == 0 {
			return fmt.Errorf("remote %s does not exist", rslug)
		}

		slug := args.Get(1)
		if err := validateSlug(slug); err != nil {
			return err
		}

		if i, _ := c.GID_RootObjects("Host", []string{slug + "." + rSuffixSlug}, false); i > 0 {
			return fmt.Errorf("identity %s already exists under remote %s", slug, rslug)
		}

		idfile := args.Get(2) // not validating bc freedom

		c.GID_RootObjectSetFirst("Host", []string{slug + "." + rSuffixSlug}, true, ssh_config.GitIDCommonChildren{
			IdentityFile: idfile, IdentitiesOnly: idfile != "",
			XGitConfigUsername: ctx.String("username"), XGitConfigUserEmail: ctx.String("email"), XGitConfigSigningKey: ctx.String("signing-key"),
			XDescription: ctx.String("description"), XParent: rSuffixSlug,
		})

		return c.Write()
	},
}

// git-id  set
var cmdSet = &cli.Command{
	Name:      "set",
	Usage:     "Set attributes of an identity",
	ArgsUsage: "git-id set <remote> <identity> [-i, --idfile IdentityFile] [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
	Flags: []cli.Flag{
		&cli.PathFlag{Name: "idfile", Aliases: []string{"i"}, Usage: "IdentityFile to set"},
		flagDesc,
		flagGitUsername, flagGitEmail, flagGitSigningkey,
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("expected exactly 2 arguments, got %d", args.Len())
		}

		c := gidOpenConfig(ctx.Path("config"))

		rslug, slug := args.Get(0), args.Get(1)
		fullSlug := fmt.Sprintf("%s.%s.%s", slug, rslug, globalTLD)

		i, trees := c.GID_RootObjects("Host", []string{fullSlug}, false)
		if i == 0 {
			return fmt.Errorf("identity %s does not exist under remote %s", slug, rslug)
		}

		tree := trees[0]

		if ctx.IsSet("idfile") {
			tree.Children.IdentityFile = ctx.String("idfile")
		}
		if ctx.IsSet("description") {
			tree.Children.XDescription = ctx.String("description")
		}
		if ctx.IsSet("username") {
			tree.Children.XGitConfigUsername = ctx.String("username")
		}
		if ctx.IsSet("email") {
			tree.Children.XGitConfigUserEmail = ctx.String("email")
		}
		if ctx.IsSet("signing-key") {
			tree.Children.XGitConfigSigningKey = ctx.String("signing-key")
		}

		c.GID_RootObjectSetFirst("Host", []string{fullSlug}, false, tree.Children)

		return c.Write()
	},
}

// git-id remove
var cmdRemove = &cli.Command{
	Name:      "remove",
	Usage:     "Remove an identity",
	ArgsUsage: "git-id remove <remote> <identity> <-y, --yes>",
	Flags: []cli.Flag{
		flagAckRemove,
		flagConfig,
	},
	Action: func(ctx *cli.Context) error {
		args := ctx.Args()
		if args.Len() != 2 {
			return fmt.Errorf("expected exactly 2 arguments, got %d", args.Len())
		}

		c := gidOpenConfig(ctx.Path("config"))

		rslug, slug := args.Get(0), args.Get(1)
		fullSlug := fmt.Sprintf("%s.%s.%s", slug, rslug, globalTLD)

		i, _ := c.GID_RootObjects("Host", []string{fullSlug}, false)
		if i == 0 {
			return fmt.Errorf("identity %s does not exist under remote %s", slug, rslug)
		}

		c.GID_RootObjectRemoveFirst("Host", []string{fullSlug})

		return c.Write()
	},
}
