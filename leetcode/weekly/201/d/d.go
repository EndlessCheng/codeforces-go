package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minCost(length int, a []int) (ans int) {
	const inf int = 1e9
	n := len(a)
	dp := make([][]int, n+5)
	for i := range dp {
		dp[i] = make([]int, n+5)
	}
	sort.Ints(a)
	a = append(append([]int{0}, a...), length)
	for sz := 1; sz <= n+1; sz++ {
		for l := 0; l+sz <= n+1; l++ {
			r := l + sz
			min := inf
			for i := l + 1; i < r; i++ {
				if v := dp[l][i] + dp[i][r] + a[r] - a[l]; v < min {
					min = v
				}
			}
			if min < inf {
				dp[l][r] = min
			}
		}
	}
	return dp[0][n+1]
}
