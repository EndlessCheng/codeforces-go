package main

// https://space.bilibili.com/206214/dynamic
func checkXMatrix(grid [][]int) bool {
	for i, row := range grid {
		for j, v := range row {
			if v == 0 == (i == j || i+j == len(grid)-1) {
				return false
			}
		}
	}
	return true
}
