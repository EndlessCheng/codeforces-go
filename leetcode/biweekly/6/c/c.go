package main

import (
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func mostVisitedPattern(username []string, timestamp []int, website []string) (ans []string) {
	type pair struct {
		ts   int
		site string
	}
	mp := map[string][]pair{}
	for i, name := range username {
		mp[name] = append(mp[name], pair{timestamp[i], website[i]})
	}
	cnt := map[[3]string]int{}
	for _, a := range mp {
		sort.Slice(a, func(i, j int) bool { return a[i].ts < a[j].ts })
		vis := map[[3]string]bool{}
		for i, v := range a {
			for j := i + 1; j < len(a); j++ {
				w := a[j]
				for _, x := range a[j+1:] {
					vis[[3]string{v.site, w.site, x.site}] = true
				}
			}
		}
		for p := range vis {
			cnt[p]++
		}
	}
	mx := 0
	for p, c := range cnt {
		if c > mx {
			mx = c
			ans = []string{p[0], p[1], p[2]}
		} else if c == mx && (p[0] < ans[0] || p[0] == ans[0] && (p[1] < ans[1] || p[1] == ans[1] && p[2] < ans[2])) {
			ans = []string{p[0], p[1], p[2]}
		}
	}
	return
}
