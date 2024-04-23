package main

// https://space.bilibili.com/206214
func findColumnWidth(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		mn, mx := 0, 0
		for _, row := range grid {
			mn = min(mn, row[j])
			mx = max(mx, row[j])
		}
		xLen := 1
		for x := max(mx/10, -mn); x > 0; x /= 10 {
			xLen++
		}
		ans[j] = max(ans[j], xLen)
	}
	return ans
}

func findColumnWidth2(grid [][]int) []int {
	ans := make([]int, len(grid[0]))
	for j := range grid[0] {
		for _, row := range grid {
			xLen := 0
			if row[j] <= 0 {
				xLen = 1
			}
			for x := row[j]; x != 0; x /= 10 {
				xLen++
			}
			ans[j] = max(ans[j], xLen)
		}
	}
	return ans
}
