package main

import "slices"

// https://space.bilibili.com/206214
func maxValue1(nums []int) int64 {
	// 先计算整个 nums 的交替和
	alterSum := 0
	for i, x := range nums {
		alterSum += x * (1 - i%2*2)
	}

	n := len(nums)
	f := make([]int, n+1)
	for i := 1; i < n; i++ {
		d := (nums[i] - nums[i-1]) * (i%2*2 - 1)
		f[i+1] = max(f[i-1], 0) + d
	}

	return int64(alterSum + slices.Max(f)*2)
}

func maxValue(nums []int) int64 {
	alterSum := nums[0]
	var f0, f1, mx int
	for i := 1; i < len(nums); i++ {
		alterSum += nums[i] * (1 - i%2*2)
		d := (nums[i] - nums[i-1]) * (i%2*2 - 1)
		f0, f1 = f1, max(f0, 0)+d
		mx = max(mx, f1)
	}
	return int64(alterSum + mx*2)
}
