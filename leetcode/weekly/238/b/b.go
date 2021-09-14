package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxFrequency(a []int, k int) (ans int) {
	sort.Ints(a)
	sum := make([]int, len(a)+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r, v := range a {
		l := sort.Search(r, func(l int) bool { return (r-l)*v-sum[r]+sum[l] <= k })
		ans = max(ans, r-l+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
