package main

import "math"

// https://space.bilibili.com/206214
func maxScore(grid [][]int) int {
	ans := math.MinInt
	colMin := make([]int, len(grid[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	for _, row := range grid {
		preMin := math.MaxInt
		for j, x := range row {
			ans = max(ans, x-min(preMin, colMin[j]))
			colMin[j] = min(colMin[j], x)
			preMin = min(preMin, colMin[j])
		}
	}
	return ans
}
