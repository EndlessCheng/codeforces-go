package main

import "slices"

// https://space.bilibili.com/206214
func minCost(nums []int) int64 {
	n := len(nums)
	a := nums[1 : n-1]
	slices.Sort(a)

	ans := int64(nums[0] + nums[n-1])
	for i := 1; i < n-2 && a[i-1]+a[i] < 0; i += 2 {
		ans += int64(a[i-1] + a[i])
	}
	return ans
}
