package main

import "sort"

// https://space.bilibili.com/206214
func sortPeople(names []string, b []int) (ans []string) {
	type pair struct{ x string; y int }
	ps := make([]pair, len(names))
	for i, v := range names {
		ps[i] = pair{v, b[i]}
	}
	sort.Slice(ps, func(i, j int) bool {
		a, b := ps[i], ps[j]
		return a.y > b.y
	})
	for _, p := range ps {
		ans = append(ans, p.x)
	}
	return
}
