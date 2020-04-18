package main

func numWays(n int, relation [][]int, k int) (ans int) {
	g := make([][]int, n)
	for _, p := range relation {
		v, w := p[0], p[1]
		g[v] = append(g[v], w)
	}
	var f func(v, dep int)
	f = func(v, dep int) {
		if dep == k {
			if v == n-1 {
				ans++
			}
			return
		}
		for _, w := range g[v] {
			f(w, dep+1)
		}
	}
	f(0, 0)
	return
}
