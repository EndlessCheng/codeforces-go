package main

func minReorder(n int, connections [][]int) (ans int) {
	type edge struct{ to, delta int }
	g := make([][]edge, n)
	for _, e := range connections {
		v, w := e[0], e[1]
		g[v] = append(g[v], edge{w, 1})
		g[w] = append(g[w], edge{v, 0})
	}
	var f func(v, fa int)
	f = func(v, fa int) {
		for _, e := range g[v] {
			if w := e.to; w != fa {
				ans += e.delta
				f(w, v)
			}
		}
	}
	f(0, -1)
	return
}
