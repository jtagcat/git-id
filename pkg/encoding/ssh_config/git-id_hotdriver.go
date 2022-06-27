package ssh_config

import "strings"

// WARN: made _only_ for git-id, may break

// package compatible with OpenSSH 8.8

type GitIDCommonChildren struct {
	IdentitiesOnly bool
	IdentityFile, Hostname, XDescription,
	XGitConfigUsername, XGitConfigUserMail, XGitConfigSigningKey,
	XParent string
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

	if c.Hostname != "" {
		raw = append(raw, RawKeyword{
			Key:    "Hostname",
			Values: []RawValue{{Value: c.Hostname, Quoted: 2}},
		})
	}
	if c.XParent != "" {
		raw = append(raw, RawKeyword{
			Key:    "XParent",
			Values: []RawValue{{Value: c.XParent, Quoted: 2}},
		})
	}

	if c.XGitConfigUsername != "" {
		raw = append(raw, RawKeyword{
			Key:    "XGitConfig",
			Values: []RawValue{{Value: "user.name"}, {Value: c.XGitConfigUsername, Quoted: 2}},
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
		switch strings.ToLower(r.Key) {
		default: // cowardly ignore and forget
		case "identitiesonly":
			switch strings.ToLower(r.Values[0].Value) {
			case "true", "yes":
				c.IdentitiesOnly = true
			}
		case "xdescription":
			for i, v := range r.Values {
				if i > 0 {
					c.XDescription += "\n"
				}
				c.XDescription += v.Value
			}
		case "identityfile":
			c.IdentityFile = r.Values[0].Value
		case "hostname":
			c.Hostname = r.Values[0].Value
		case "xparent":
			c.XParent = r.Values[0].Value
		case "xgitconfig":
			switch strings.ToLower(r.Values[0].Value) {
			case "user.name":
				c.XGitConfigUsername = r.Values[1].Value
			case "user.email":
				c.XGitConfigUserMail = r.Values[1].Value
			case "user.signingkey":
				c.XGitConfigSigningKey = r.Values[1].Value
			}
		}
	}

	return c
}
