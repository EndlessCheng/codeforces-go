package main

// github.com/EndlessCheng/codeforces-go
func restoreArray(es [][]int) (ans []int) {
	g := map[int][]int{}
	for _, e := range es {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var root int
	for v, vs := range g {
		if len(vs) == 1 {
			root = v
			break
		}
	}
	var f func(v, fa int)
	f = func(v, fa int) {
		ans = append(ans, v)
		for _, w := range g[v] {
			if w != fa {
				f(w, v)
			}
		}
	}
	f(root, 1e9)
	return
}
