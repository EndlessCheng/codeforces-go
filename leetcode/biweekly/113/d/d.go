package main

// https://space.bilibili.com/206214
func minEdgeReversals(n int, edges [][]int) (ans []int) {
	type pair struct{ to, dir int }
	g := make([][]pair, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], pair{y, 1})
		g[y] = append(g[y], pair{x, -1})
	}

	ans = make([]int, n)
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, e := range g[x] {
			y := e.to
			if y != fa {
				if e.dir < 0 {
					ans[0]++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(x, fa int) {
		for _, e := range g[x] {
			y := e.to
			if y != fa {
				ans[y] = ans[x] + e.dir
				reroot(y, x)
			}
		}
	}
	reroot(0, -1)
	return ans
}
