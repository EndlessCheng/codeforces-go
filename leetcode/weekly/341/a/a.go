package main

// https://space.bilibili.com/206214
func rowAndMaximumOnes(mat [][]int) []int {
	maxSum, idx := -1, 0
	for i, row := range mat {
		sum := 0
		for _, x := range row {
			sum += x
		}
		if sum > maxSum {
			maxSum, idx = sum, i
		}
	}
	return []int{idx, maxSum}
}
