package main

// github.com/EndlessCheng/codeforces-go
func tree3(e []int) int {
	g := make([][]int, len(e)+2)
	for w, v := range e {
		w += 2
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

	isEnd := make([]bool, len(g))
	f = func(v, fa, d int) {
		if d == maxD {
			isEnd[v] = true
			return
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(dv, 0, 0)
	f(dw, 0, 0)

	cntEnds := 0
	for _, is := range isEnd {
		if is {
			cntEnds++
		}
	}
	if cntEnds > 2 {
		return maxD
	}
	return maxD - 1
}
