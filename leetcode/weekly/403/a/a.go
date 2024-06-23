package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	ans := math.MaxInt
	for i, n := 0, len(nums); i < n/2; i++ {
		ans = min(ans, nums[i]+nums[n-1-i])
	}
	return float64(ans) / 2
}
