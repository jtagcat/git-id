package ssh_config

import "strings"

// WARN: made _only_ for git-id, may break
func (c *Config) GID_InsertRootComment(s string) {
	c.cfg = append(c.cfg, RawTopLevel{Comment: s})
}

func (c *Config) GID_PreappendInclude(i string) {
	c.cfg = append(
		[]RawTopLevel{{Key: "Include", Values: []RawValue{{Value: i, Quoted: 2}}}},
		c.cfg...)
}

//
//

// WARN: made _only_ for git-id, may break
func (c *Config) GID_RootObjectCount(key string, firstValue string) (matches int) {
	for _, root := range c.cfg {
		if strings.EqualFold(root.Key, key) &&
			len(root.Values) < 2 &&
			strings.EqualFold(root.Values[0].Value, firstValue) {

			matches++
		}
	}
	return matches
}

// WARN: made _only_

// WARN: made _only_ for git-id, may break
// Brutal implementation of override and don't care anything for the sake of time
func (c *Config) GIDRootObjectSet(key string, values []string, childs GitIDCommonChildren) {
	children := childsEncode(childs)
	for _, root := range c.cfg {
		if strings.EqualFold(root.Key, key) {
			var valuesComparable []string
			for _, v := range root.Values {
				valuesComparable = append(valuesComparable, v.Value)
			}

			if EqualFoldSlice(valuesComparable, values) {
				root.Children = children
				return
			}
		}
	}
	var valuesInjectable []RawValue
	for _, v := range values {
		valuesInjectable = append(valuesInjectable, RawValue{Value: v})
	}

	c.cfg = append(c.cfg, RawTopLevel{
		Key:      key,
		Values:   valuesInjectable,
		Children: children,
	})
}

// WARN: made _only_ for git-id, may break
func (c *Config) GIDRootObjectExists(key, firstValue string) (ok bool, secondValue string) {
	for _, root := range c.cfg {
		if strings.EqualFold(root.Key, key) {
			if strings.EqualFold(root.Values[0].Value, firstValue) {
				return true, root.Values[1].Value
			}
		}
	}
	return false, ""
}

func (c *Config) GIDRootObjectRemove(key string, values []string) (ok bool) {
	i := func(config []RawTopLevel) int {
		for i, root := range config {
			if strings.EqualFold(root.Key, key) {
				var valuesComparable []string
				for _, v := range root.Values {
					valuesComparable = append(valuesComparable, v.Value)
				}

				if EqualFoldSlice(valuesComparable, values) {
					return i
				}
			}
		}
		return -1
	}(c.cfg)
	if i == -1 {
		return false
	}

	c.cfg = append(c.cfg[:i], c.cfg[i+1:]...)
	return true
}

//
// Match OriginalHost github.com
//   IdentityFile ~/.ssh/id_rsa
//
// Host *.gh.git-id
//   Hostname github.com
//   #XDescription "iz GitHub"
//
// Host jc.gh.git-id
//  IdentityFile ~/.ssh/id_rsa # this is redundant with defaults, IdentityFile is used for matching the default to an ident
//  #XGitConfig user.name jtagcat
//  #XGitConfig user.email blah
//  #XDescription uwu
// Host w.gh.git-id
//  IdentityFile ~/.ssh/work_sk
//
