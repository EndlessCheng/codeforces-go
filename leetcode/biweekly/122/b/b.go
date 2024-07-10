package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
func canSortArray(nums []int) bool {
	preMax := 0
	for i, n := 0, len(nums); i < n; {
		mx := 0
		ones := bits.OnesCount(uint(nums[i]))
		for ; i < n && bits.OnesCount(uint(nums[i])) == ones; i++ {
			if nums[i] < preMax { // 无法排成有序的
				return false
			}
			mx = max(mx, nums[i]) // 更新本组最大值
		}
		preMax = mx
	}
	return true
}
