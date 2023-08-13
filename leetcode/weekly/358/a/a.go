package main

import "math"

// https://space.bilibili.com/206214
func maxSum(nums []int) int {
	ans := -1
	maxVal := [10]int{}
	for i := range maxVal {
		maxVal[i] = math.MinInt // 表示不存在最大值
	}
	for _, v := range nums {
		maxD := 0
		for x := v; x > 0; x /= 10 {
			maxD = max(maxD, x%10)
		}
		ans = max(ans, v+maxVal[maxD])
		maxVal[maxD] = max(maxVal[maxD], v)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
