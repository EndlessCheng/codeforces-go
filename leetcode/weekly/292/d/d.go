package main

// github.com/EndlessCheng/codeforces-go
func hasValidPath(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	if (m+n)%2 == 0 || grid[0][0] == ')' || grid[m-1][n-1] == '(' { // 剪枝
		return false
	}

	vis := make([][][]bool, m)
	for i := range vis {
		vis[i] = make([][]bool, n)
		for j := range vis[i] {
			vis[i][j] = make([]bool, (m+n+1)/2)
		}
	}
	var dfs func(x, y, c int) bool
	dfs = func(x, y, c int) bool {
		if c > m-x+n-y-1 { // 剪枝：即使后面都是 ')' 也不能将 c 减为 0
			return false
		}
		if x == m-1 && y == n-1 { // 终点
			return c == 1 // 终点一定是 ')'
		}
		if vis[x][y][c] { // 重复访问
			return false
		}
		vis[x][y][c] = true
		if grid[x][y] == '(' {
			c++
		} else if c--; c < 0 { // 非法括号字符串
			return false
		}
		return x < m-1 && dfs(x+1, y, c) || y < n-1 && dfs(x, y+1, c) // 往下或者往右
	}
	return dfs(0, 0, 0) // 起点
}
