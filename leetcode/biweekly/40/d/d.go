package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func lisAll(a []int) []int {
	n := len(a)
	lis := make([]int, n)
	dp := make([]int, 0, n)
	for i, v := range a {
		if p := sort.SearchInts(dp, v); p < len(dp) {
			dp[p] = v
			lis[i] = p + 1
		} else {
			dp = append(dp, v)
			lis[i] = len(dp)
		}
	}
	return lis
}

func minimumMountainRemovals(a []int) (ans int) {
	pre := lisAll(a)
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
	suf := lisAll(a)
	for i := 1; i < n-1; i++ {
		if pre[i] > 1 && suf[n-1-i] > 1 {
			ans = max(ans, pre[i]+suf[n-1-i]-1)
		}
	}
	return n - ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
