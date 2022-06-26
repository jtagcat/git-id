package ssh_config

import "strings"

// WARN: made _only_ for git-id, may break

// package compatible with OpenSSH 8.8

type GitIDCommonChildren struct {
	IdentitiesOnly bool
	XDescription, IdentityFile,
	XGitConfigUserName, XGitConfigUserMail, XGitConfigSigningKey string
}

// WARN: made _only_ for git-id, may break
func childsEncode(c GitIDCommonChildren) (raw []RawKeyword) {
	// for-range??? reflection??
	if c.IdentitiesOnly {
		raw = append(raw, RawKeyword{
			Key:    "IdentitiesOnly",
			Values: []RawValue{{Value: "true", Quoted: 2}},
		})
	}

	if c.XDescription != "" {
		var values []RawValue
		for _, s := range strings.Split(c.XDescription, "\n") {
			values = append(values, RawValue{Value: s, Quoted: 2})
		}
		raw = append(raw, RawKeyword{
			Key:    "XDescription",
			Values: values,
		})
	}

	if c.IdentityFile != "" {
		raw = append(raw, RawKeyword{
			Key:    "IdentityFile",
			Values: []RawValue{{Value: c.IdentityFile, Quoted: 2}},
		})
	}

	if c.XGitConfigUserName != "" {
		raw = append(raw, RawKeyword{
			Key:    "XGitConfig",
			Values: []RawValue{{Value: "user.name"}, {Value: c.XGitConfigUserName, Quoted: 2}},
		})
	}

	if c.XGitConfigUserMail != "" {
		raw = append(raw, RawKeyword{
			Key:    "XGitConfig",
			Values: []RawValue{{Value: "user.email"}, {Value: c.XGitConfigUserMail, Quoted: 2}},
		})
	}

	if c.XGitConfigSigningKey != "" {
		raw = append(raw, RawKeyword{
			Key:    "XGitConfig",
			Values: []RawValue{{Value: "user.signingkey"}, {Value: c.XGitConfigSigningKey, Quoted: 2}},
		})
	}

	return raw
}

func childsDecode(raw []RawKeyword) (c GitIDCommonChildren) {
	// expecting panics from out of index, but that's ok (:
	for _, r := range raw {
		switch r.Key {
		default: // cowardly ignore and forget
		case "IdentitiesOnly":
			switch r.Values[0].Value {
			case "true", "yes":
				c.IdentitiesOnly = true
			}
		case "XDescription":
			for i, v := range r.Values {
				if i > 0 {
					c.XDescription += "\n"
				}
				c.XDescription += v.Value
			}
		case "IdentityFile":
			c.IdentityFile = r.Values[0].Value
		case "XGitConfig":
			switch r.Values[0].Value {
			case "user.name":
				c.XGitConfigUserName = r.Values[1].Value
			case "user.email":
				c.XGitConfigUserMail = r.Values[1].Value
			case "user.signingkey":
				c.XGitConfigSigningKey = r.Values[1].Value
			}
		}
	}

	return c
}
