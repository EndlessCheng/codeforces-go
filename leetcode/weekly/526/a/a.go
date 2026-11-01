package main

import (
	"math"
)

// https://space.bilibili.com/206214
func minMergedInteger(nums []int) int {
	ans := math.MaxInt
	for i := 0; i < len(nums); i += 2 {
		x, y := nums[i], nums[i+1]
		for v := nums[i+1]; v > 0; v /= 10 {
			x *= 10
		}
		ans = min(ans, x+y)
	}
	return ans
}
