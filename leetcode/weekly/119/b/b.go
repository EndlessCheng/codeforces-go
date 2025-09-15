package main

import "slices"

// https://space.bilibili.com/206214
func largestPerimeter(nums []int) int {
	slices.Sort(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0 // 无解
}
