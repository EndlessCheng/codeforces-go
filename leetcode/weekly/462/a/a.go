package main

// https://space.bilibili.com/206214
func reverseSubmatrix(grid [][]int, x, y, k int) [][]int {
	l, r := x, x+k-1
	for l < r {
		for j := y; j < y+k; j++ {
			grid[l][j], grid[r][j] = grid[r][j], grid[l][j]
		}
		l++
		r--
	}
	return grid
}
