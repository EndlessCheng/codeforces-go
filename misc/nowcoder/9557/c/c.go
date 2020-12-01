package main

// github.com/EndlessCheng/codeforces-go
func tree3(e []int) int {
	n := len(e) + 2
	g := make([][]int, n)
	for w, v := range e {
		w += 2
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	mx, u, c := -1, 0, make([]int, n)
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > mx {
			mx, u = d, v
		}
		c[d]++
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(1, 0, 0)
	mx, c = -1, make([]int, n)
	f(u, 0, 0)
	f(u, 0, 0)
	if c[mx] == 2 {
		mx--
	}
	return mx
}
