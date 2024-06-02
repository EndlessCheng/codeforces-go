package main

import "math"

// https://space.bilibili.com/206214
func minimumDifference(nums []int, k int) int {
	ans := math.MaxInt
	for i, x := range nums {
		ans = min(ans, abs(x-k))
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
			ans = min(ans, abs(nums[j]-k))
		}
	}
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
