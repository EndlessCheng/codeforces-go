package main

import "math"

// https://space.bilibili.com/206214
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		cnt, unique, mn := make([]int, n), 0, math.MaxInt
		for j := i; j >= 0; j-- {
			x := nums[j]
			cnt[x]++
			if cnt[x] == 1 { // 首次出现
				unique++
			} else if cnt[x] == 2 { // 不再唯一
				unique--
			}
			mn = min(mn, f[j]-unique)
		}
		f[i+1] = mn + k
	}
	return f[n] + n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
