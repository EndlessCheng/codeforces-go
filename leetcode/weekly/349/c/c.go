package main

import "slices"

// https://space.bilibili.com/206214
func minCost(nums []int, x int) int64 {
	n := len(nums)
	s := make([]int64, n)
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			mn = min(mn, nums[j%n]) // 从 nums[i] 到 nums[j%n] 的最小值
			s[j-i] += int64(mn)     // 累加操作 j-i 次的花费
		}
	}
	return slices.Min(s)
}
