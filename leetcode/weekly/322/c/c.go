package main

import "math"

// https://space.bilibili.com/206214
func minScore(n int, roads [][]int) int {
	type edge struct{ to, dis int }
	g := make([][]edge, n+1)
	for _, e := range roads {
		x, y, dis := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, dis})
		g[y] = append(g[y], edge{x, dis})
	}

	vis := make([]bool, n+1)
	ans := math.MaxInt

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true // 避免重复访问
		for _, e := range g[x] {
			ans = min(ans, e.dis)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}

	// 遍历节点 1 所在连通块
	dfs(1)
	return ans
}
