package main

import "slices"

// https://space.bilibili.com/206214
func maximumOr(nums []int, k int) int64 {
	allOr, multi := 0, 0
	for _, x := range nums {
		multi |= allOr & x
		allOr |= x
	}

	ans := 0
	for _, x := range nums {
		ans = max(ans, x<<k|(allOr^x)|multi)
	}
	return int64(ans)
}

func maximumOr2(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n+1)
	for i, x := range slices.Backward(nums) {
		suf[i] = suf[i+1] | x
	}

	ans, pre := 0, 0
	for i, x := range nums {
		ans = max(ans, pre|x<<k|suf[i+1])
		pre |= x
	}
	return int64(ans)
}
