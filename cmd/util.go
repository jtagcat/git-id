package cmd

import (
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"go.uber.org/zap"
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

func gidOpenConfig(name string) *ssh_config.Config {
	c, new, err := ssh_config.OpenConfig(ssh_config.Opts{
		SubXKeys: []string{
			"XDescription",
			"XGitConfig",
		},
		Indent: "  ",
	}, name)
	if err != nil {
		zap.L().Fatal("couldnt open config", zap.String("path", name), zap.Error(err))
	}

	if !new {
		return c
	}

	// init
	c.GID_InsertRootComment(gitidHeaderInfo)
	c.Write() // before including
	zap.L().Info("created config file", zap.String("path", name))

	// import
	u, _, err := ssh_config.OpenConfig(ssh_config.Opts{Indent: "  "}, userSSHConfigFile)
	if err != nil {
		zap.L().Fatal("couldn't open user config", zap.String("path", userSSHConfigFile))
	}

	// search
	if i, _ := u.GID_RootObjectCount("import", []string{name}, false); i == 0 {
		u.GID_PreappendInclude(name) // ew
		u.Write()
	}
	zap.L().Info("included config in ssh_config", zap.String("path", userSSHConfigFile))

	return c
}
