package main

// https://space.bilibili.com/206214
func minimumArea(grid [][]int) int {
	left, right := len(grid[0]), 0
	top, bottom := len(grid), 0
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				left = min(left, j)
				right = max(right, j)
				top = min(top, i)
				bottom = i
			}
		}
	}
	return (right - left + 1) * (bottom - top + 1)
}
