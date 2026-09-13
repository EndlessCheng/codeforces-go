package main

// https://space.bilibili.com/206214
func cyclicShift(n int, grid [][]int, rowShift, colShift []int) [][]int {
	for i, row := range grid {
		shift := rowShift[i]
		grid[i] = append(row[shift:], row[:shift]...)
	}

	column := make([]int, n)
	for j, shift := range colShift {
		// 收集列元素
		col := column[:0]
		for _, row := range grid[shift:] {
			col = append(col, row[j])
		}
		for _, row := range grid[:shift] {
			col = append(col, row[j])
		}
		// 填入列
		for i, row := range grid {
			row[j] = col[i]
		}
	}
	return grid
}
