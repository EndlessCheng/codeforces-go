package main

// https://space.bilibili.com/206214
func countIslands(grid [][]int, k int) (ans int) {
	dirs := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		res := grid[i][j]
		grid[i][j] = 0 // 标记 (i,j) 访问过
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 {
				res += dfs(x, y)
			}
		}
		return res
	}
	for i, row := range grid {
		for j, x := range row {
			if x > 0 && dfs(i, j)%k == 0 {
				ans++
			}
		}
	}
	return
}
