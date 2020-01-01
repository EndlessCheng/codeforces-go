package main

import "sort"

func smallestRangeII(a []int, k int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(a)
	sort.Ints(a)
	ans := a[n-1] - a[0]
	for i := 1; i < n; i++ {
		up := max(a[i-1]+k, a[n-1]-k)
		down := min(a[i]-k, a[0]+k)
		ans = min(ans, up-down)
	}
	return ans
}
