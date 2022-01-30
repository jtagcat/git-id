package ssh_config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessXComments(t *testing.T) {
	want := []RawTopLevel{
		{Comment: " This file is managed by git-id"},
		{Key: "XHeader", Values: []RawValue{{"Identities", 0}}},
		{Key: "Host", Values: []RawValue{{"jc.gh.git-id", 0}}, Children: []RawKeyword{
			{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/id_rsa", 0}}},
			{Key: "XGitConfig", Values: []RawValue{{"user.name", 0}, {"jtagcat", 0}}, Comment: " it is me!"},
			{},
			{Comment: " Random comment"},
			{Key: "XGitConfig", Values: []RawValue{{"user.email", 0}, {"user@domain.tld", 0}}},
			{Key: "XDescription", Values: []RawValue{{"uwu", 0}}},
		}},
		{},
		{Key: "Host", Values: []RawValue{{"foo.gh.git-id", 0}, {"foo.sh.git-id", 1}}, EncodingKVSeperatorIsEquals: true, Children: []RawKeyword{
			{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/foo_sk", 0}}},
		}},
		{Key: "XHeader", Values: []RawValue{{"Remotes", 0}}},
		{},
		{Key: "Host", Values: []RawValue{{"*.gh.git-id", 0}}, Children: []RawKeyword{
			{Key: "HostName", EncodingKVSeperatorIsEquals: true, Values: []RawValue{{"github.com", 0}}},
			{Key: "XDescription", Values: []RawValue{{"iz GitHub", 2}}},
			{Key: "IdentitiesOnly", Values: []RawValue{{"yes", 0}}},
		}},
		{Key: "Host", Values: []RawValue{{"*.sh.git-id", 0}}, Children: []RawKeyword{
			{Key: "Hostname", Values: []RawValue{{"git.sr.ht", 0}}},
			{Comment: " Child comment"},
		}},
		{},
		{Comment: " Root comment"},
	}

	cfg, _ := DecodeToRaw(exampleConfig())
	got, err := DecodeXKeys(cfg, "x")
	assert.Nil(t, err)
	fmt.Printf("%v", got)
	assert.Equal(t, got, want)
}
