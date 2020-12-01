package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func tree3(e []int) int {
	g := make([][]int, len(e)+2)
	for w, v := range e {
		w += 2
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var maxD, u int
	ds := []int{}
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		ds = append(ds, d)
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	maxD = -1
	f(1, 0, 0)
	maxD, ds = -1, nil
	f(u, 0, 0)
	f(u, 0, 0)
	sort.Ints(ds)
	return ds[len(ds)-3]
}
