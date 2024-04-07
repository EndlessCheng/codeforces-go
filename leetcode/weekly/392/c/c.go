package main

import "slices"

// https://space.bilibili.com/206214
func minOperationsToMakeMedianK(nums []int, k int) (ans int64) {
	slices.Sort(nums)
	m := len(nums) / 2
	if nums[m] > k {
		for i := m; i >= 0 && nums[i] > k; i-- {
			ans += int64(nums[i] - k)
		}
	} else {
		for i := m; i < len(nums) && nums[i] < k; i++ {
			ans += int64(k - nums[i])
		}
	}
	return
}
