package main

import (
	"math"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func minimumDifference(nums []int, k int) int {
	slices.Sort(nums)
	ans := math.MaxInt
	for i := k - 1; i < len(nums); i++ {
		ans = min(ans, nums[i]-nums[i-k+1])
	}
	return ans
}
