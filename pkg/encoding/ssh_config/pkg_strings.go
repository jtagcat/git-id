package ssh_config

// in the interest of not confusing the lang server,
// we aren't splitting it to pkg (can't put it in parent pkg either)

import (
	"sort"
	"strings"
)

// strings.Cut, but with multiple seperators, found is either empty or seperator
func CutAny(s, chars string) (before, after string, found string) {
	if i := strings.IndexAny(s, chars); i >= 0 {
		return s[:i], s[i+1:], string(s[i])
	}
	return s, "", ""
}

func EqualFoldSlice(s []string, t []string) bool {
	if len(s) != len(t) {
		return false
	}

	sort.SliceStable(s, func(i, j int) bool { return s[i] < s[j] })
	sort.SliceStable(t, func(i, j int) bool { return s[i] < s[j] })

	for i := range s {
		if !strings.EqualFold(s[i], t[i]) {
			return false
		}
	}
	return true
}
