package ssh_config

import "strings"

// WARN: made _only_ for git-id, may break
func (c *Config) GID_PreappendRootComment(s string) {
	c.cfg = append([]RawTopLevel{{Comment: s}}, c.cfg...)
}

// WARN: made _only_ for git-id, may break
func (c *Config) GID_PreappendInclude(i string) {
	c.cfg = append(
		[]RawTopLevel{{Key: "Include", Values: []RawValue{{Value: i, Quoted: 2}}}},
		c.cfg...)
}

type miniTree struct {
	Values   []string
	Children GitIDCommonChildren
}

// WARN: made _only_ for git-id, may break
// suffix: value is handled as a suffix
func (c *Config) GID_RootObjectCount(key string, values []string, wildcard bool) (matches int, trees []miniTree) {
	for _, root := range c.cfg {
		if strings.EqualFold(root.Key, key) &&
			valuesMatch(root.Values, values, wildcard) {

			var foundValues []string
			for _, v := range root.Values {
				foundValues = append(foundValues, v.Value)
			}

			trees = append(trees, miniTree{
				Values:   foundValues,
				Children: childsDecode(root.Children),
			})
			matches++
		}
	}
	return matches, trees
}

// order and len() matters, "" means ignore
// suffix: value is handled as a suffix
func valuesMatch(against []RawValue, values []string, suffix bool) bool {
	if len(against) != len(values) {
		return false
	}

	for i, a := range against {
		if values[i] != "" {
			if suffix {
				if !strings.HasSuffix(strings.ToLower(a.Value), strings.ToLower(values[i])) {
					return false
				}
			}

			if !strings.EqualFold(a.Value, values[i]) {
				return false
			}
		}
	}
	return true
}

// WARN: made _only_ for git-id, may break
// Brutal implementation of override and don't care anything for the sake of time
func (c *Config) GID_RootObjectSetFirst(key string, values []string, firstValueIsSubKey bool, childs GitIDCommonChildren) {
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
func (c *Config) GIDRootObjectRemoveFirst(key string, values []string) (ok bool) {
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
