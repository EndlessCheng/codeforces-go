package main

// https://space.bilibili.com/206214
func numberOfRightTriangles(grid [][]int) (ans int64) {
	n := len(grid[0])
	colSum := make([]int, n)
	for j := 0; j < n; j++ {
		for _, row := range grid {
			colSum[j] += row[j]
		}
	}
	for _, row := range grid {
		rowSum := -1
		for _, x := range row {
			rowSum += x
		}
		for j, x := range row {
			if x == 1 {
				ans += int64(rowSum * (colSum[j] - 1))
			}
		}
	}
	return
}
