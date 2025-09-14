package main

import "slices"

// https://space.bilibili.com/206214
func maxKDistinct(nums []int, k int) []int {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	nums = slices.Compact(nums)
	if len(nums) > k {
		return nums[:k]
	}
	return nums
}
