package main

func minTime(n int, edges [][]int, has []bool) (ans int) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(v, fa, d int) bool
	f = func(v, fa, d int) (found bool) {
		if has[v] {
			ans += d
			found = true
			d = 0
		}
		for _, w := range g[v] {
			if w != fa {
				if f(w, v, d+1) {
					found = true
					d = 0
				}
			}
		}
		if found && v > 0 {
			ans++
		}
		return
	}
	f(0, -1, 0)
	return
}
