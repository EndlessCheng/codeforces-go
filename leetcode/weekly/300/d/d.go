package main

// https://space.bilibili.com/206214/dynamic
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func countPaths(grid [][]int) (ans int) {
	const mod = 1_000_000_007
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := 1
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j] {
				res = (res + dfs(x, y)) % mod
			}
		}
		*p = res
		return res
	}
	for i, row := range grid {
		for j := range row {
			ans += dfs(i, j)
		}
	}
	return ans % mod
}
