package main

import (
	//. "nc_tools"
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// github.com/EndlessCheng/codeforces-go
func solve(n int, x int, Edge []*Point) int {
	if x == 1 {
		return 0
	}
	g := make([][]int, n+1)
	for _, e := range Edge {
		v, w := e.X, e.Y
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	d1 := make([]int, n+1)
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		d1[v] = d
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(1, -1, 0)
	ans := 0
	f = func(v, fa, d int) {
		dd := d1[v]
		if d <= dd && dd > ans {
			ans = dd
		}
		if d >= dd {
			return
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(x, -1, 0)
	return ans + 1
}
