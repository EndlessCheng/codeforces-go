package main

// https://space.bilibili.com/206214
func maxMoves(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		ans = max(ans, j)
		if ans == n-1 { // ans 已达到最大值
			return
		}
		// 向右上/右/右下走一步
		for k := max(i-1, 0); k < min(i+2, m); k++ {
			if grid[k][j+1] > grid[i][j] {
				dfs(k, j+1)
			}
		}
		grid[i][j] = 0
	}
	for i := range grid {
		dfs(i, 0) // 从第一列的任一单元格出发
	}
	return
}
