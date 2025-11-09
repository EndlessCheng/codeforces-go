package main

import "math"

// https://space.bilibili.com/206214
// 3573. 买卖股票的最佳时机 V
func maximumProfit(prices []int, l, r, k int) int64 {
	n := len(prices)
	f := make([][3]int, k+2)
	for j := 1; j <= k+1; j++ {
		f[j][1] = math.MinInt / 2
		f[j][2] = math.MinInt / 2
	}
	f[0][0] = math.MinInt / 2
	for i := l; i < r; i++ {
		p := prices[i%n]
		for j := k + 1; j > 0; j-- {
			f[j][0] = max(f[j][0], f[j][1]+p, f[j][2]-p)
			f[j][1] = max(f[j][1], f[j-1][0]-p)
			f[j][2] = max(f[j][2], f[j-1][0]+p)
		}
	}
	return int64(f[k+1][0])
}

func maximumScore(nums []int, k int) int64 {
	n := len(nums)
	maxI := 0
	for i, x := range nums {
		if x > nums[maxI] {
			maxI = i
		}
	}

	ans1 := maximumProfit(nums, maxI, maxI+n, k)
	ans2 := maximumProfit(nums, maxI+1, maxI+1+n, k)
	return max(ans1, ans2)
}
