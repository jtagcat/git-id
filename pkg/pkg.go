package pkg

import (
	"strings"
)

func EqualFoldSlice(s []string, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if !strings.EqualFold(s[i], t[i]) {
			return false
		}
	}
	return true
}

//TODO:
// func FlushToSlice(a &[]generic, i &generic) {
// a = append(a, i)
// i = nil
// }
