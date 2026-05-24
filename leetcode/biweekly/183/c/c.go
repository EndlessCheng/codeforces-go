package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
// 53. 最大子数组和（子数组长度 >= 2）
func maxSubArray(nums []int) int {
	ans := math.MinInt // 注意答案可以是负数，不能初始化成 0
	f := nums[0]
	for _, x := range nums[1:] {
		ans = max(ans, f+x) // f+x 保证子数组至少有两个数
		f = max(f, 0) + x
	}
	return ans
}

func maxScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := math.MinInt

	// 单独计算子数组长为 1 的情况，此时子数组不能在矩形边界上
	if m > 2 && n > 2 {
		for _, row := range grid[1 : m-1] {
			ans = max(ans, slices.Max(row[1:n-1]))
		}
	}

	// 每行的最大子数组和（子数组长度 >= 2）
	for _, row := range grid {
		ans = max(ans, maxSubArray(row))
	}

	// 每列的最大子数组和（子数组长度 >= 2）
	col := make([]int, m)
	for j := range n {
		for i, row := range grid {
			col[i] = row[j]
		}
		ans = max(ans, maxSubArray(col))
	}

	return ans
}
