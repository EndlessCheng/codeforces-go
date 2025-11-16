package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maximizeExpressionOfThree1(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	return nums[n-1] + nums[n-2] - nums[0]
}

func maximizeExpressionOfThree(nums []int) int {
	mx, mx2, mn := math.MinInt, math.MinInt, math.MaxInt
	for _, x := range nums {
		if x > mx {
			mx2 = mx
			mx = x
		} else if x > mx2 {
			mx2 = x
		}
		mn = min(mn, x)
	}
	return mx + mx2 - mn
}
