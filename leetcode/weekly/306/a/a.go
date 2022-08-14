package main

// https://space.bilibili.com/206214
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	for i, row := range grid[:n-2] {
		for j := 0; j < n-2; j++ {
			mx := 0
			for _, r := range grid[i : i+3] {
				for _, v := range r[j : j+3] {
					mx = max(mx, v)
				}
			}
			row[j] = mx
		}
		grid[i] = row[:n-2]
	}
	return grid[:n-2]
}

func max(a, b int) int { if b > a { return b }; return a}
