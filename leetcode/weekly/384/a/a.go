package main

// https://space.bilibili.com/206214
func modifiedMatrix(matrix [][]int) [][]int {
	for j := range matrix[0] {
		mx := 0
		for _, row := range matrix {
			mx = max(mx, row[j])
		}
		for _, row := range matrix {
			if row[j] == -1 {
				row[j] = mx
			}
		}
	}
	return matrix
}
