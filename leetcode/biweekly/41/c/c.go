package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func stoneGameVI(a, b []int) (ans int) {
	type pair struct{ x, y int }
	ps := make([]pair, len(a))
	for i, v := range a {
		ps[i] = pair{v, b[i]}
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x+a.y > b.x+b.y })
	for i, p := range ps {
		if i&1 > 0 {
			ans -= p.y
		} else {
			ans += p.x
		}
	}
	if ans > 0 {
		ans = 1
	} else if ans < 0 {
		ans = -1
	}
	return
}
