package main

import "math"

// https://space.bilibili.com/206214
func minCost(nums []int, x int) int64 {
	n := len(nums)
	sum := make([]int, n)
	for i := range sum {
		sum[i] += i * x // 操作 i 次
	}
	for i, mn := range nums {
		for j := i; j < n+i; j++ {
			mn = min(mn, nums[j%n])
			sum[j-i] += mn // 操作 j-i 次
		}
	}
	ans := math.MaxInt
	for _, s := range sum {
		ans = min(ans, s)
	}
	return int64(ans)
}

func min(a, b int) int { if b < a { return b }; return a }
