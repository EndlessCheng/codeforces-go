package main

// github.com/EndlessCheng/codeforces-go
func PointsOnDiameter(n int, vs, ws []int) (ans int) {
	g := make([][]int, n+1)
	for i, v := range vs {
		w := ws[i]
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
	f(1, 0, 0)
	dv := u
	maxD = -1
	f(u, 0, 0)
	dw := u

	onDiameter := make([]bool, len(g))
	var f2 func(v, fa, d int) bool
	f2 = func(v, fa, d int) bool {
		if d == maxD {
			onDiameter[v] = true
			return true
		}
		for _, w := range g[v] {
			if w != fa {
				if f2(w, v, d+1) {
					onDiameter[v] = true
				}
			}
		}
		return onDiameter[v]
	}
	f2(dv, 0, 0)
	f2(dw, 0, 0)
	for _, on := range onDiameter {
		if on {
			ans++
		}
	}
	return
}
