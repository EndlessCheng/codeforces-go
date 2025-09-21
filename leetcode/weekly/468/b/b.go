package main

import "slices"

// https://space.bilibili.com/206214
func maxTotalValue(nums []int, k int) int64 {
	return int64(slices.Max(nums)-slices.Min(nums)) * int64(k)
}
