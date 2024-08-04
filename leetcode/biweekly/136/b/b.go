package main

// https://space.bilibili.com/206214
func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	diffRow := 0
	for _, row := range grid {
		for j := 0; j < n/2; j++ {
			if row[j] != row[n-1-j] {
				diffRow++
			}
		}
	}

	diffCol := 0
	for j := 0; j < n; j++ {
		for i, row := range grid[:m/2] {
			if row[j] != grid[m-1-i][j] {
				diffCol++
			}
		}
	}

	return min(diffRow, diffCol)
}
