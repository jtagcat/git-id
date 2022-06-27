package cmd

import "github.com/urfave/cli/v2"

// git-id how
var cmdHow = &cli.Command{
	Name:      "how",
	Usage:     "Describe how git-id works",
	ArgsUsage: "git-id how",
	Flags:     []cli.Flag{},
	Hidden:    true,
	// -t OR -u OR both: test
	// -u: get non-default user
	// id: use that OR get from pwd / -C
	// ...
}
