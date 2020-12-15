package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func solve(n int, a []int) int {
	sort.Ints(a)
	var mi, mx int
	for i := n - 1; i > 1; i-- {
		if a[i] < a[i-1]+a[i-2] {
			mx = a[i] + a[i-1] + a[i-2]
			break
		}
	}
	for i := 1; i < n-1; i++ {
		if j := sort.SearchInts(a[:i], a[i+1]-a[i]+1); j < i {
			mi = a[j] + a[i] + a[i+1]
			break
		}
	}
	// 题目数据保证能找到
	return mx - mi
}
