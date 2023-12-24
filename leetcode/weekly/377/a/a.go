package main

import "slices"

// https://space.bilibili.com/206214
func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}
