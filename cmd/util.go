package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	spkg "github.com/jtagcat/git-id/pkg/encoding/ssh_config/pkg"

	valid "github.com/asaskevich/govalidator"
	"github.com/urfave/cli/v2"
)

// 	hostParts := strings.Split(host, ".")
// 	tld := hostParts[len(hostParts)-1]
// 	if tld == flGI_TLD { // iz already git-id
// 		if len(hostParts) != 3 {
// 			return nil, fmt.Errorf("%w: host has git-id TLD, but is does not have 3 parts", fs.ErrInvalid)
// 		}
// 		// we have user.remote.git-id, forward remote
// 		// TODO: potential bugish behaviour: from x.a.git-id to y.b.git-id, where b and a have the same underlieing domain
// 		// var a: intended to see them as different (preferred)
// 		// var b: get remote Hostname, and do same as non-hijcacked hosts
// 		return []string{hostParts[1]}, nil
// 	}

// 	// non-hijacked host
// 	gitidRemotes := ssh_config.RootBySubKV(cfg, "Host", "Hostname", []string{host})
// 	if len(gitidRemotes) == 0 {
// 		log.Fatal().Str("host", host).Msg("no git-id remote for host found")
// 	}
// 	// we got a list of "*.xyz.git-id"
// 	// extract git-id remote and forward
// 	for _, idr := range gitidRemotes {
// 		for _, v := range idr.Values {
// 			giHostParts := strings.Split(v.Value, ".")
// 			if len(giHostParts) == 3 && giHostParts[0] == "*" && giHostParts[2] == flGI_TLD {
// 				remotes = append(remotes, giHostParts[1])
// 			}
// 		}
// 	}
// 	return remotes, nil
// }

// uses globals: flConfigPath, gitidHeaderInfo
// func openConfig() (*os.File, error) {
// 	var err error

// 	flConfigPath, err = homedir.Expand(flConfigPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if _, err := os.Stat(name); err == nil {
// 		return os.OpenFile(name, flag, 0000)
// 	}
// 	return nil, fs.ErrNotExist

// 	f, err := pkg.OpenFileExisting(flConfigPath, os.O_RDONLY)
// 	if err != nil {
// 		return f, nil
// 	}

// 	// Init config

// 	os.WriteFile()

// 	_, err = f.Write([]byte(
// 		ssh_config.EncodeFromRaw([]ssh_conf)
// 		// gitidHeaderInfo+\n
// 		))
// 	return err

// 	// Import to .ssh/config
// 	// if already exists
// }

func gidOpenConfig(path string) *ssh_config.Config {
	c, new, err := ssh_config.OpenConfig(ssh_config.Opts{
		SubXKeys: []string{
			"XDescription",
			"XGitConfig",
			"XParent",
		},
		Indent: "  ",
	}, path)
	if err != nil {
		log.Fatalf("couldnt open config at %s: %e", path, err)
	}

	if !new {
		return c
	}

	// init
	c.GID_PreappendRootComment(gitidHeaderInfo)
	c.Write() // before including
	log.Printf("created config file at %s", path)

	// import
	u, _, err := ssh_config.OpenConfig(ssh_config.Opts{Indent: "  "}, userSSHConfigFile)
	if err != nil {
		log.Fatalf("couldnt open user config at %s: %e", userSSHConfigFile, err)
	}

	// search
	if i, _ := u.GID_RootObjects("import", []string{path}, false); i == 0 {
		u.GID_PreappendInclude(path) // ew
		u.Write()
	}
	log.Printf("included config in ssh_config at %s", path)

	return c
}

func inInvalids(name string) (in bool) {
	for _, i := range invalidNames {
		if strings.EqualFold(i, name) {
			return true
		}
	}
	return false
}

var invalidNames []string

func init() {
	invalidNames = getinvalidsRoot(App)
}

func getinvalidsRoot(a *cli.App) (i []string) {
	for _, c := range a.Commands {
		i = append(i, getInvalids(c.Subcommands)...)
	}
	return i
}

func getInvalids(cmds []*cli.Command) (i []string) {
	for _, c := range cmds {
		i = append(append(append(i,

			c.Names()...), c.Aliases...),

			getInvalids(c.Subcommands)...)
	}
	return i
}

// potentially buggy if used unwell
func remoteSlug(fullSlug string) string {
	if strings.Count(fullSlug, ".") < 2 {
		return ""
	}

	s := strings.TrimPrefix(fullSlug, "*.")
	s, _, _ = spkg.CutLast(s, ".")

	return s
}

func identSlug(fullSlug, remoteSlug string) string {
	i := strings.Index(fullSlug, remoteSlug)
	return fullSlug[:i]
}

func validateSlug(slug string) error {
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

	return nil
}
