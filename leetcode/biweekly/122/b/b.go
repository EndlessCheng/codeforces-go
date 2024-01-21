package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
func canSortArray(nums []int) bool {
	preMax := 0
	for i, n := 0, len(nums); i < n; {
		mn, mx := nums[i], nums[i]
		ones := bits.OnesCount(uint(mn))
		for i++; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			mn = min(mn, nums[i])
			mx = max(mx, nums[i])
		}
		if mn < preMax {
			return false
		}
		preMax = mx
	}
	return true
}
