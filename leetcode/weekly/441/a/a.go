package main

import (
	"math"
)

// https://space.bilibili.com/206214
func maxSum(nums []int) (ans int) {
	mx := math.MinInt
	set := map[int]struct{}{}
	for _, x := range nums {
		if x < 0 {
			mx = max(mx, x)
		} else if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			ans += x
		}
	}
	if len(set) == 0 {
		return mx
	}
	return
}
