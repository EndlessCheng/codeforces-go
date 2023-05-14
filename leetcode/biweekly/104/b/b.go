package main

import "sort"

// https://space.bilibili.com/206214
func matrixSum(nums [][]int) (ans int) {
	for _, row := range nums {
		sort.Ints(row)
	}
	for j := range nums[0] {
		mx := 0
		for _, row := range nums {
			mx = max(mx, row[j])
		}
		ans += mx
	}
	return
}

func max(a, b int) int { if a < b { return b }; return a }
