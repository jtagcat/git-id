package cmd

import (
	"fmt"
	"strings"

	"github.com/gravitational/teleport/lib/asciitable"
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
				tree.Children.XGitConfigUsername, tree.Children.XGitConfigUserMail, tree.Children.XGitConfigSigningKey,
				tree.Children.XDescription,
			})
		}
	}

	t.DiscardEmpty()
	fmt.Println(t.AsBuffer().String())
	return nil
}

// git-id add
var cmdAdd = &cli.Command{
	Name:      "add",
	Usage:     "Add an identity",
	ArgsUsage: "git-id add <remote> <identity> <IdentityFile> [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
	Flags: []cli.Flag{
		flagDesc,
		&cli.StringFlag{Name: "username", Aliases: []string{"u"}, Usage: "git config user.name"},
		&cli.StringFlag{Name: "email", Aliases: []string{"e"}, Usage: "git config user.email"},
		&cli.StringFlag{Name: "signing-key", Aliases: []string{"sk"}, Usage: "git config user.signingKey"},
		flagConfig,
	},
	Hidden: true,
}

// git-id  set
var cmdSet = &cli.Command{
	Name:      "set",
	Usage:     "Add an identity",
	ArgsUsage: "git-id set <remote> <identity> [-i IdentityFile] [-d, --description] [-u, --username] [-e, --email] [-sk, --signing-key]",
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
}

// git-id remove
var cmdRemove = &cli.Command{
	Name:      "remove",
	Usage:     "Remove an identity",
	ArgsUsage: "git-id remove <remote> <identity> <-y, --yes>",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "yes", Aliases: []string{"-y"}, Usage: "acknowledge potential breakage"},
		flagConfig,
	},
	Hidden: true,
}
