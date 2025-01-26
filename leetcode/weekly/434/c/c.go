package main

import "math"

func maxFrequency(nums []int, k int) int {
	f0 := 0
	f1 := [51]int{}
	f2 := math.MinInt
	maxF1 := math.MinInt
	for _, x := range nums {
		f2 = max(f2, maxF1)
		f1[x] = max(f1[x], f0) + 1
		if x == k {
			f2++
			f0++
		}
		maxF1 = max(maxF1, f1[x])
	}
	return max(maxF1, f2)
}

func maxFrequency1(nums []int, k int) (ans int) {
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}

	for target := range set {
		f0, f1, f2 := 0, math.MinInt, math.MinInt
		for _, x := range nums {
			f2 = max(f2, f1) + b2i(x == k)
			f1 = max(f1, f0) + b2i(x == target)
			f0 += b2i(x == k)
		}
		ans = max(ans, f1, f2)
	}
	return
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
