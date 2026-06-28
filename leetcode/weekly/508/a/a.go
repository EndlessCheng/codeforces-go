package main

import "slices"

// https://space.bilibili.com/206214
func maxSum(nums []int, k int, mul int) (ans int64) {
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	for _, x := range nums[:k] {
		ans += int64(x) * int64(max(mul, 1))
		mul--
	}
	return
}
