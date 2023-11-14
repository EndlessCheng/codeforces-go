package main

import "slices"

// https://space.bilibili.com/206214
func maximizeSum(nums []int, k int) int {
	return (slices.Max(nums)*2 + k - 1) * k / 2
}
