package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minOperations(nums []int) int {
	n := len(nums)
	if n < 3 {
		return -1
	}
	mx := slices.Max(nums) + n - 1

	dp := func() [][]int {
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, mx)
		}

		for j := 1; j < mx; j++ {
			f[0][j] = abs(nums[0] - j)
		}
		for i := 1; i < n-1; i++ {
			preMin := math.MaxInt
			for j := i + 1; j < mx; j++ {
				preMin = min(preMin, f[i-1][j-1])
				f[i][j] = preMin + abs(nums[i]-j)
			}
		}

		return f
	}

	pre := dp()
	slices.Reverse(nums)
	suf := dp()
	slices.Reverse(suf)
	slices.Reverse(nums)

	// 枚举峰顶的位置和值
	ans := math.MaxInt
	for i := 1; i < n-1; i++ {
		// 从左到右严格递增，j 至少是 i+1
		// 从右到左严格递增，j 至少是 n-i
		for j := max(i+1, n-i); j < mx; j++ {
			ans = min(ans, pre[i][j]+suf[i][j]-abs(nums[i]-j))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
