package main

// https://space.bilibili.com/206214
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findMaxFish(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 0
		}
		sum := grid[x][y]
		grid[x][y] = 0 // 标记成访问过
		for _, d := range dirs { // 四方向移动
			sum += dfs(x+d.x, y+d.y)
		}
		return sum
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
