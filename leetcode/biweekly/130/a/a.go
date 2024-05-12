package main

// https://space.bilibili.com/206214
func satisfiesConditions(grid [][]int) bool {
	for i, row := range grid {
		for j, x := range row {
			if j > 0 && x == row[j-1] || i > 0 && x != grid[i-1][j] {
				return false
			}
		}
	}
	return true
}
