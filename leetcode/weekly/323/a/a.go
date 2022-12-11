package main

import "sort"

// https://space.bilibili.com/206214
func deleteGreatestValue(grid [][]int) (ans int) {
	for _, row := range grid {
		sort.Ints(row)
	}
	for j := range grid[0] {
		mx := 0
		for _, row := range grid {
			mx = max(mx, row[j])
		}
		ans += mx
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
