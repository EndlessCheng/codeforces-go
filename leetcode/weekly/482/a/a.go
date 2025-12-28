package main

import "math"

// https://space.bilibili.com/206214
func maximumScore1(nums []int) int64 {
	n := len(nums)
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], nums[i])
	}

	ans := math.MinInt
	preSum := 0
	for i, x := range nums[:n-1] {
		preSum += x
		ans = max(ans, preSum-sufMin[i+1])
	}
	return int64(ans)
}

func maximumScore(nums []int) int64 {
	preSum := 0
	for _, x := range nums {
		preSum += x
	}

	ans := math.MinInt
	sufMin := math.MaxInt
	for i := len(nums) - 1; i > 0; i-- {
		preSum -= nums[i] // 撤销
		sufMin = min(sufMin, nums[i])
		ans = max(ans, preSum-sufMin)
	}
	return int64(ans)
}
