package ssh_config

import (
	"strings"
)

func exampleConfig2() *strings.Reader {
	return strings.NewReader(
		"# This file is managed by git-id\n" + // 1
			"#XHeader Identities\n" + // 2
			"Host jc.gh.git-id\n" + // 3
			"  IdentityFile ~/.ssh/id_rsa\n" + // 4
			"  #XGitConfig user.name jtagcat # it is me!\n" + // 5
			// no 6
			"# Random comment\n" + // 7
			"  #XGitConfig user.email user@domain.tld\n" + // 8
			"  #XDescription uwu\n" + // 9
			"\n" + // 10
			"Host=foo.gh.git-id 'foo.sh.git-id'\n" + // 11
			"IdentityFile ~/.ssh/foo_sk\n" + // 12
			"#XHeader Remotes\n" + // 13
			"\n" + // 14
			"Host *.gh.git-id\n" + // 15
			"  HostName=github.com\n" + // 16
			"  #XDescription \"iz GitHub\" \n" + // 17
			"  IdentitiesOnly yes\n" + // 18
			"Host *.sh.git-id\n" + // 19
			"  Hostname git.sr.ht\n" + // 20
			"  # Child comment\n" + // 21
			"\n" + // 22
			"  # Root comment") // 23
}
