package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func findValueOfPartition(nums []int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	for i := 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-1])
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
