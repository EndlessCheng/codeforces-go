package main

import (
	"sort"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func beforeAndAfterPuzzles(a []string) []string {
	mp := map[string]bool{}
	for i, v := range a {
		sp := strings.Split(v, " ")
		for j, w := range a {
			if j == i {
				continue
			}
			sp2 := strings.Split(w, " ")
			if sp2[0] == sp[len(sp)-1] {
				mp[strings.Join(append(sp, sp2[1:]...), " ")] = true
			}
		}
	}
	keys := make([]string, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
