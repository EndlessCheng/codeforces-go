package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minOperations(a []int, x int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	if sum[n] < x {
		return -1
	}
	ans = 1e9
	for i := n; i >= 0; i-- {
		v := sum[n] - sum[i]
		j := sort.SearchInts(sum[:i+1], x-v)
		if sum[j]+v == x && j+n-i < ans {
			ans = j + n - i
		}
	}
	if ans < 1e9 {
		return
	}
	return -1
}
