package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minPairSum(a []int) (ans int) {
	sort.Ints(a)
	n := len(a)
	for i, v := range a[:n/2] {
		ans = max(ans, v+a[n-1-i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
