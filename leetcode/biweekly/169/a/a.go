package main

import "slices"

// https://space.bilibili.com/206214
func minMoves(nums []int) int {
	ans := slices.Max(nums) * len(nums)
	for _, x := range nums {
		ans -= x
	}
	return ans
}
