package main

import "math"

// https://space.bilibili.com/206214
func findMissingElements(nums []int) (ans []int) {
	mn, mx := math.MaxInt, math.MinInt
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
		mn = min(mn, x)
		mx = max(mx, x)
	}

	for i := mn + 1; i < mx; i++ {
		if !has[i] {
			ans = append(ans, i)
		}
	}
	return
}
