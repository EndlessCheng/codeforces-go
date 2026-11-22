package main

import "math"

// https://space.bilibili.com/206214
func minCost(nums []int) int64 {
	n := len(nums)
	ans := math.MaxInt
	start := 0
	for i, x := range nums {
		if i == n-1 || x != nums[i+1] { // i 是连续相同段的末尾
			if i > 0 && start < n-1 { // 这一段不能只有 nums[0] 或者只有 nums[n-1]
				ans = min(ans, (start+n-i-1)*x)
			}
			start = i + 1 // 下一段的开始位置
		}
	}
	return int64(ans)
}
