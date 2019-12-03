package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type pair struct {
		v      int
		isBlue bool
		length int
	}
	g := make([][]pair, n)
	for _, p := range redEdges {
		g[p[0]] = append(g[p[0]], pair{p[1], false, 0})
	}
	for _, p := range blueEdges {
		g[p[0]] = append(g[p[0]], pair{p[1], true, 0})
	}

	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = -1
	}
	visB := make([]bool, n)
	visR := make([]bool, n)
	q := []pair{{}}
	for len(q) > 0 {
		var p pair
		p, q = q[0], q[1:]
		for _, e := range g[p.v] {
			w, isBlue := e.v, e.isBlue
			if p.v > 0 && isBlue == p.isBlue {
				continue
			}
			if isBlue && visB[w] || !isBlue && visR[w] {
				continue
			}
			if isBlue {
				visB[w] = true
			} else {
				visR[w] = true
			}
			if ans[w] == -1 {
				ans[w] = p.length + 1
			} else {
				ans[w] = min(ans[w], p.length+1)
			}
			q = append(q, pair{w, isBlue, p.length + 1})
		}
	}
	return ans
}
