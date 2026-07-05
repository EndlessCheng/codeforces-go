package main

import "slices"

// https://space.bilibili.com/206214
func isMiddleElementUnique(nums []int) bool {
	m := len(nums) / 2
	return !slices.Contains(nums[:m], nums[m]) &&
		!slices.Contains(nums[m+1:], nums[m])
}
