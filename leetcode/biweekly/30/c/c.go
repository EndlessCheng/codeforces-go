package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minDifference(a []int) (ans int) {
	if len(a) < 5 {
		return
	}
	sort.Ints(a)
	ans = int(2e9)
	for i := 0; i < 4; i++ {
		ans = min(ans, a[len(a)-4+i]-a[i])
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
