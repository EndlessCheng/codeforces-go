package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func smallestRangeII(a []int, k int) (ans int) {
	n := len(a)
	sort.Ints(a)
	ans = a[n-1] - a[0]
	for i := 0; i < n-1; i++ {
		ans = min(ans, max(a[i]+k, a[n-1]-k)-min(a[0]+k, a[i+1]-k))
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
