package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func minOperations(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	nums = slices.Compact(nums) // 原地去重

	ans, left := 0, 0
	for i, x := range nums {
		for nums[left] < x-n+1 { // nums[left] 不在窗口中
			left++
		}
		ans = max(ans, i-left+1)
	}
	return n - ans
}
