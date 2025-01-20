package main

import "slices"

// https://space.bilibili.com/206214
func minCost(a, b []int, k int64) int64 {
	ans2 := int64(0)
	for i, x := range a {
		ans2 += int64(abs(x - b[i]))
	}
	if ans2 <= k {
		return ans2
	}

	slices.Sort(a)
	slices.Sort(b)
	ans1 := k
	for i, x := range a {
		ans1 += int64(abs(x - b[i]))
	}

	return min(ans1, ans2)
}

func abs(x int) int { if x < 0 { return -x }; return x }
