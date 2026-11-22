package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minCost1(nums []int) int64 {
	n := len(nums)
	a := nums[1 : n-1]
	slices.Sort(a)

	ans := int64(nums[0] + nums[n-1])
	for i := 1; i < n-2 && a[i-1]+a[i] < 0; i += 2 {
		ans += int64(a[i-1] + a[i])
	}
	return ans
}

func minCost2(nums []int) int64 {
	n := len(nums)
	ans := nums[0] + nums[n-1]
	neg := 0
	maxNeg := math.MinInt32
	minNonNeg := math.MaxInt32

	for _, x := range nums[1 : n-1] {
		if x < 0 {
			ans += x
			neg++
			maxNeg = max(maxNeg, x)
		} else {
			minNonNeg = min(minNonNeg, x)
		}
	}

	if neg%2 > 0 {
		if maxNeg+minNonNeg < 0 {
			ans += minNonNeg
		} else {
			ans -= maxNeg
		}
	}

	return int64(ans)
}

func minCost(nums []int) int64 {
	n := len(nums)
	f0, f1 := 0, math.MaxInt/2
	for _, x := range nums[1 : n-1] {
		f0, f1 = min(f0, f1+x), min(f1, f0+x)
	}
	return int64(nums[0] + nums[n-1] + f0)
}
