package main

// github.com/EndlessCheng/codeforces-go
func countSubgraphsForEachDiameter(n int, edges [][]int) (ans []int) {
	f := func(sub int) (res int) {
		g := make([][]int, n)
		st := 0
		for i, e := range edges {
			if sub>>i&1 > 0 {
				v, w := e[0]-1, e[1]-1
				st = v
				g[v] = append(g[v], w)
				g[w] = append(g[w], v)
			}
		}

		vis := make([]bool, n)
		var maxD, u int
		var f func(v, d int)
		f = func(v, d int) {
			vis[v] = true
			if d > maxD {
				maxD, u = d, v
			}
			for _, w := range g[v] {
				if !vis[w] {
					f(w, d+1)
				}
			}
		}
		maxD = -1
		f(st, 0)
		for i, b := range vis {
			if !b && len(g[i]) > 0 {
				return
			}
		}
		vis = make([]bool, n)
		maxD = -1
		f(u, 0)
		return maxD
	}
	ans = make([]int, n)
	for sub := 0; sub < 1<<len(edges); sub++ {
		ans[f(sub)]++
	}
	return ans[1:]
}
