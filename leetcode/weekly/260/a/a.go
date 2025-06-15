package main

import "math"

// github.com/EndlessCheng/codeforces-go
func maximumDifference(nums []int) (ans int) {
	preMin := math.MaxInt
	for _, x := range nums {
		ans = max(ans, x-preMin) // 把 x 当作 nums[j]
		preMin = min(preMin, x)  // 把 x 当作 nums[i]
	}
	if ans == 0 {
		ans = -1
	}
	return
}
