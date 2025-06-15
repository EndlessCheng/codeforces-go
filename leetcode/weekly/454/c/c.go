package main

import "math"

// https://space.bilibili.com/206214
func maximumProduct(nums []int, m int) int64 {
	ans := math.MinInt
	mn, mx := math.MaxInt, math.MinInt
	for i := m - 1; i < len(nums); i++ {
		// 维护 [0,i-m+1] 中的最小值和最大值
		y := nums[i-m+1]
		mn = min(mn, y)
		mx = max(mx, y)
		// 枚举右
		x := nums[i]
		ans = max(ans, x*mn, x*mx)
	}
	return int64(ans)
}
