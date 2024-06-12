package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func findValueOfPartition(nums []int) int {
	slices.Sort(nums)
	ans := math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}
