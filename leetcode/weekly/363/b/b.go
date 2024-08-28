package main

import "slices"

// https://space.bilibili.com/206214
func countWays(nums []int) (ans int) {
	slices.Sort(nums)
	if nums[0] > 0 {
		ans = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < i && i < nums[i] {
			ans++
		}
	}
	return ans + 1
}
