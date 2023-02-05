package main

// https://space.bilibili.com/206214
func isPossibleToCutPath(g [][]int) bool {
	m, n := len(g), len(g[0])

	var dfs func(int, int) bool
	dfs = func(x, y int) bool { // 返回能否到达终点
		if x == m-1 && y == n-1 {
			return true
		}
		g[x][y] = 0 // 直接修改
		return x < m-1 && g[x+1][y] > 0 && dfs(x+1, y) ||
			  y < n-1 && g[x][y+1] > 0 && dfs(x, y+1)
	}

	// 提前特判一些可以直接得到答案的情况
	return m*n > 2 && (m == 1 || n == 1 || g[1][0] == 0 || g[0][1] == 0 || g[m-2][n-1] == 0 || g[m-1][n-2] == 0 ||
		   !dfs(1, 0) || !dfs(0, 1))
}
