package main

func makeConnected(n int, edges [][]int) int {
	if len(edges)+1 < n {
		return -1
	}
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	vis := make([]bool, n)
	cnt := 0
	var f func(int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !vis[i] {
			cnt++
			f(i)
		}
	}
	return cnt - 1
}
