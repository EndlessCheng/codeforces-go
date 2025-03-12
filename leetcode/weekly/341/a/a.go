package main

// https://space.bilibili.com/206214
func rowAndMaximumOnes(mat [][]int) []int {
	rowIdx, maxSum := 0, -1
	for i, row := range mat {
		s := 0
		for _, x := range row {
			s += x
		}
		if s > maxSum {
			rowIdx, maxSum = i, s
		}
	}
	return []int{rowIdx, maxSum}
}
