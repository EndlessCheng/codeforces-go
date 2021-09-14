package main

func frogPosition(n int, edges [][]int, time int, target int) (ans float64) {
	g := make([][]int, n)
	g[0] = append(g[0], -1)
	for _, e := range edges {
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	target--
	var f func(v, fa, t int, p float64)
	f = func(v, fa, t int, p float64) {
		// corner case 需要细心
		if v == target && (t == 0 || t > 0 && len(g[v]) == 1) {
			ans = p
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, t-1, p/float64(len(g[v])-1))
			}
		}
	}
	f(0, -1, time, 1)
	return
}
