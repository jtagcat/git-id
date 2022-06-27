package pkg

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

// strings.Cut, but starting from last character, found is either empty or seperator
func CutLast(s, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
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
