package main

// https://space.bilibili.com/206214
func numberOfPaths(grid [][]int, k int) int {
	const mod int = 1e9 + 7
	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, k)
		}
	}
	f[0][1][0] = 1
	for i, row := range grid {
		for j, x := range row {
			for v := 0; v < k; v++ {
				f[i+1][j+1][(v+x)%k] = (f[i+1][j][v] + f[i][j+1][v]) % mod
			}
		}
	}
	return f[m][n][0]
}
