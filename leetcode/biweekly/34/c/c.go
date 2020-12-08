package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func findLengthOfShortestSubarray(a []int) (ans int) {
	if sort.IntsAreSorted(a) {
		return
	}
	n := len(a)
	l, r := 0, n-1
	for ; a[l] <= a[l+1]; l++ {
	}
	for ; a[r] >= a[r-1]; r-- {
	}
	ans = min(n-1-l, r)
	for ; r < n; r++ {
		ans = min(ans, r-sort.SearchInts(a[:l+1], a[r]+1))
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
