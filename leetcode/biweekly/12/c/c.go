package main

// github.com/EndlessCheng/codeforces-go
func treeDiameter(es [][]int) (ans int) {
	g := make([][]int, len(es)+1)
	for _, e := range es {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	maxD, u := -1, 0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > maxD {
			maxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	maxD = -1
	f(u, -1, 0)
	return maxD
}
