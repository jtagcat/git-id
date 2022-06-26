package ssh_config

// in the interest of not confusing the lang server,
// we aren't splitting it to pkg (can't put it in parent pkg either)

import "strings"

// strings.Cut, but with multiple seperators, found is either empty or seperator
func CutAny(s, chars string) (before, after string, found string) {
	if i := strings.IndexAny(s, chars); i >= 0 {
		return s[:i], s[i+1:], string(s[i])
	}
	return s, "", ""
}
