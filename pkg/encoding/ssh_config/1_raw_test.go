package ssh_config

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func exampleConfig() (Opts, *strings.Reader) {
	return Opts{
			RootXKeys: map[string]bool{
				"xheader": false,
				"rooter":  true,
			},
			SubXKeys: []string{"XGitConfig", "XDescription"},
			Indent:   "  ",
		},
		strings.NewReader(
			"# This file is managed by git-id\n" + // 1
				"#XHeader Identities\n" + // 2
				"#rOoTeR\n" +
				"  #XDescription\n" +
				"\n" +
				"#rooter two\n" +
				"Include rootlevel\n" +
				"Host jc.gh.git-id\n" +
				"  IdentityFile ~/.ssh/id_rsa\n" + // 5
				"  #XGitConfig user.name jtagcat # it is me!\n" +
				"\n" +
				"  # Random comment\n" +
				"  #XGitConfig user.email user@domain.tld\n" +
				"  #XDescription uwu\n" + // 10
				"\n" +
				"Host=foo.gh.git-id 'foo.sh.git-id'\n" +
				"  IdentityFile ~/.ssh/foo_sk\n" +
				"  Include sublevel\n" +
				"#XHeader Remotes\n" + // 15
				"\n" +
				"Host *.gh.git-id\n" +
				"  HostName=github.com\n" +
				"  #XDescription \"iz GitHub\"\n" +
				"  IdentitiesOnly yes\n" + // 20
				"  # XDescription this is not a key this is comment\n" +
				"Host *.sh.git-id\n" +
				"  Hostname git.sr.ht\n" +
				"  # Child comment\n" +
				"\n" + // 25
				"# Root comment" +
				"\n") // 27
}

func TestDecodeToRaw(t *testing.T) {
	want := []RawTopLevel{
		{Comment: " This file is managed by git-id"},
		{Key: "XHeader", Values: []RawValue{{"Identities", 0}}},
		{Key: "rOoTeR", Children: []RawKeyword{
			{Key: "XDescription"},
		}},
		{},
		{Key: "rooter", Values: []RawValue{{"two", 0}}},
		{Key: "Include", Values: []RawValue{{"rootlevel", 0}}},
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
			{Key: "Include", Values: []RawValue{{"sublevel", 0}}},
		}},
		{Key: "XHeader", Values: []RawValue{{"Remotes", 0}}},
		{},
		{Key: "Host", Values: []RawValue{{"*.gh.git-id", 0}}, Children: []RawKeyword{
			{Key: "HostName", EncodingKVSeperatorIsEquals: true, Values: []RawValue{{"github.com", 0}}},
			{Key: "XDescription", Values: []RawValue{{"iz GitHub", 2}}},
			{Key: "IdentitiesOnly", Values: []RawValue{{"yes", 0}}},
			{Comment: " XDescription this is not a key this is comment"},
		}},
		{Key: "Host", Values: []RawValue{{"*.sh.git-id", 0}}, Children: []RawKeyword{
			{Key: "Hostname", Values: []RawValue{{"git.sr.ht", 0}}},
			{Comment: " Child comment"},
		}},
		{},
		{Comment: " Root comment"},
	}

	got, err := Decode(exampleConfig())
	assert.Nil(t, err)
	fmt.Printf("%v", got)
	assert.Equal(t, want, got)
}

func TestEncodeRawRoundtrip(t *testing.T) {
	o, want := exampleConfig()

	cfg, err := Decode(exampleConfig())
	assert.Nil(t, err)

	var got bytes.Buffer
	err = Encode(o, cfg, &got)
	assert.Nil(t, err)
	assert.Equal(t, want, strings.NewReader(got.String()))
}

func TestBadConfigUnderX(t *testing.T) {
	o, want := Opts{
		RootXKeys: map[string]bool{
			"xheader": true,
		},
		Indent: "  ",
	},
		strings.NewReader(
			"Host hello\n"+
				"  IdentityFile test\n"+
				"\n"+
				"#XHeader\n"+
				"  IdentityFile invalid\n")

	_, err := Decode(o, want)
	assert.ErrorIs(t, err, ErrValidSubkeyAfterXRoot)
}
