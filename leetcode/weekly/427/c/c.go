package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maxSubarraySum1(nums []int, k int) int64 {
	minS := make([]int, k)
	for i := range k - 1 {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	s := 0
	for j, x := range nums {
		s += x
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}

func maxSubarraySum2(nums []int, k int) int64 {
	sum := make([]int, len(nums)+1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}

	minS := make([]int, k)
	for i := range minS {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	for j, s := range sum {
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}

func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	f := make([]int, n+1)
	for i := range f {
		f[i] = math.MinInt
	}

	sum := 0 // 滑动窗口维护长为 k 的子数组的元素和
	for i, x := range nums {
		sum += x
		left := i - k + 1
		if left < 0 {
			continue
		}
		f[i+1] = max(f[left], 0) + sum
		sum -= nums[left]
	}
	return int64(slices.Max(f))
}
