package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minDistance(a []int, k int) int {
	sort.Ints(a)
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(l, left int) (res int) {
		if l+left >= n {
			return
		}
		if left == 0 {
			return 1e9
		}
		dv := &dp[l][left]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for r := l + 1; r <= n; r++ {
			m := (l + r) / 2
			res = min(res, sum[r]-sum[m]-(r-m)*a[m]+(m-l)*a[m]-(sum[m]-sum[l])+f(r, left-1))
		}
		return
	}
	return f(0, k)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
