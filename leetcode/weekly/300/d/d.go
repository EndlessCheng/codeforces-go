package main

// https://space.bilibili.com/206214/dynamic
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func countPaths(grid [][]int) (ans int) {
	const mod int = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if f[i][j] != -1 {
			return f[i][j]
		}
		res := 1
		for _, d := range dirs {
			if x, y := i+d.x, j+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > grid[i][j] {
				res = (res + dfs(x, y)) % mod
			}
		}
		f[i][j] = res
		return res
	}
	for i, row := range grid {
		for j := range row {
			ans = (ans + dfs(i, j)) % mod
		}
	}
	return
}
