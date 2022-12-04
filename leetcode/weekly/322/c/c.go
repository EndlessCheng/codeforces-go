package main

import "math"

// https://space.bilibili.com/206214
func minScore(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n)
	for _, e := range roads {
		x, y, d := e[0]-1, e[1]-1, e[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}
	ans := math.MaxInt32
	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.d)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(0)
	return ans
}

func min(a, b int) int { if a > b { return b }; return a }
