package main

import (
	//. "nc_tools"
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// github.com/EndlessCheng/codeforces-go
func solve(n int, edges []*Point, a []int) int64 {
	s := 0
	type neighbor struct{ to, wt int }
	g := make([][]neighbor, n+1)
	for i, e := range edges {
		v, w, wt := e.X, e.Y, a[i]
		s += wt
		g[v] = append(g[v], neighbor{w, wt})
		g[w] = append(g[w], neighbor{v, wt})
	}
	s <<= 1

	maxD, u := -1, 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, e := range g[v] {
			if e.to != fa {
				f(e.to, v, d+e.wt)
			}
		}
	}
	f(1, 0, 0)
	maxD = -1
	f(u, 0, 0)
	return int64(s - maxD)
}
