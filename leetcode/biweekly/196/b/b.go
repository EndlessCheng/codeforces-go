package main

import "slices"

// https://space.bilibili.com/206214
func minRange(a, b, c int) int {
	nums := []int{a, b, c}
	slices.Sort(nums)
	return min(nums[2]-nums[0], nums[1])
}
