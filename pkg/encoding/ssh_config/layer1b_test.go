package ssh_config

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func exampleConfig() *strings.Reader {
	return strings.NewReader(
		"# This file is managed by git-id\n" + // 1
			"#XHeader Identities\n" + // 2
			"Host jc.gh.git-id\n" + // 3
			"  IdentityFile ~/.ssh/id_rsa\n" + // 4
			"  #XGitConfig user.name jtagcat # it is me!\n" + // 5
			"\n" + // 6
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

func TestDecodeToRaw(t *testing.T) {
	want := []RawTopLevel{
		{Comment: " This file is managed by git-id"},
		{Comment: "XHeader Identities"},
		{Key: "Host", Values: []RawValue{{"jc.gh.git-id", 0}}, Children: []RawKeyword{
			{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/id_rsa", 0}}},
			{Comment: "XGitConfig user.name jtagcat # it is me!"}}},
		{},
		{Comment: " Random comment"},
		{Comment: "XGitConfig user.email user@domain.tld"},
		{Comment: "XDescription uwu"},
		{},
		{Key: "Host", Values: []RawValue{{"foo.gh.git-id", 0}, {"foo.sh.git-id", 1}}, EncodingKVSeperatorIsEquals: true, Children: []RawKeyword{
			{Key: "IdentityFile", Values: []RawValue{{"~/.ssh/foo_sk", 0}}},
			{Comment: "XHeader Remotes"}}},
		{},
		{Key: "Host", Values: []RawValue{{"*.gh.git-id", 0}}, Children: []RawKeyword{{Key: "HostName", EncodingKVSeperatorIsEquals: true, Values: []RawValue{{"github.com", 0}}},
			{Comment: "XDescription \"iz GitHub\""}, {Key: "IdentitiesOnly", Values: []RawValue{{"yes", 0}}}}},
		{Key: "Host", Values: []RawValue{{"*.sh.git-id", 0}}, Children: []RawKeyword{
			{Key: "Hostname", Values: []RawValue{{"git.sr.ht", 0}}},
			{Comment: " Child comment"}}},
		{},
		{Comment: " Root comment"},
	}

	got, err := DecodeToRaw(exampleConfig())
	assert.Nil(t, err)
	fmt.Printf("%v", got)
	assert.Equal(t, got, want)
}
