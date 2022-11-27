package main

// https://space.bilibili.com/206214
func onesMinusZeros(grid [][]int) [][]int {
	r := make([]int, len(grid))
	c := make([]int, len(grid[0]))
	for i, row := range grid {
		for j, x := range row {
			r[i] += x*2 - 1
			c[j] += x*2 - 1
		}
	}
	for i, x := range r {
		for j, y := range c {
			grid[i][j] = x + y
		}
	}
	return grid
}
