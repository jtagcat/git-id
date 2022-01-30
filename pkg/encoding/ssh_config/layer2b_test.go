package ssh_config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessCommentPositions(t *testing.T) {
	want := []RawTopLevel{
		{Comment: " This file is managed by git-id"},
		{Comment: "XHeader Identities"},
		{Key: "Host", Values: []RawValue{{"jc.gh.git-id", 0}}, Children: []RawKeyword{{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/id_rsa", 0}}}}},
		{Comment: "XGitConfig user.name jtagcat # it is me!"},
		{Comment: " Random comment"},
		{Comment: "XGitConfig user.email user@domain.tld"},
		{Comment: "XDescription uwu"},
		{Key: "Host", Values: []RawValue{{"foo.gh.git-id", 0}, {"foo.sh.git-id", 1}},
			EncodingKVSeperatorIsEquals: true, Children: []RawKeyword{{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/foo_sk", 0}}}}},
		{Comment: "XHeader Remotes"},
		{Key: "Host", Values: []RawValue{{"*.gh.git-id", 0}}, Children: []RawKeyword{{Key: "HostName", EncodingKVSeperatorIsEquals: true, Values: []RawValue{{"github.com", 0}}}}},
		{Comment: "XDescription \"iz GitHub\"", Children: []RawKeyword{{Key: "IdentitiesOnly", Values: []RawValue{{"yes", 0}}}}},
		{Key: "Host", Values: []RawValue{{"*.sh.git-id", 0}}, Children: []RawKeyword{{Key: "Hostname", Values: []RawValue{{"git.sr.ht", 0}}}}},
	}

	got, err := DecodeToRaw(exampleConfig())
	assert.Nil(t, err)
	fmt.Printf("%v", got)
	assert.Equal(t, got, want)
}
