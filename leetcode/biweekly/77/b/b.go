package main

import "math"

// github.com/EndlessCheng/codeforces-go
func minimumAverageDifference(nums []int) (ans int) {
	pre, suf, n := 0, 0, len(nums)
	for _, v := range nums { suf += v } // 后缀和
	minDiff := math.MaxInt64
	for i, v := range nums[:n-1] {
		pre += v // 前缀和
		suf -= v // 后缀和
		d := abs(pre/(i+1) - suf/(n-1-i))
		if d < minDiff {
			minDiff, ans = d, i
		}
	}
	if (pre+nums[n-1])/n < minDiff { ans = n - 1 }
	return ans
}

func abs(x int) int { if x < 0 { return -x }; return x }
