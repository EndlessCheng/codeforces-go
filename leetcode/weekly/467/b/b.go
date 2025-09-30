package main

import "slices"

// https://space.bilibili.com/206214
func maxKDistinct(nums []int, k int) []int {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	nums = slices.Compact(nums) // 原地去重
	return nums[:min(k, len(nums))]
}
