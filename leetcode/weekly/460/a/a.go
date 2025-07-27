package main

import "slices"

// https://space.bilibili.com/206214
func maximumMedianSum(nums []int) (ans int64) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	m := len(nums) / 3
	for i := 1; i < m*2; i += 2 {
		ans += int64(nums[i])
	}
	return
}
