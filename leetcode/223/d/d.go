package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minimumTimeRequired(a []int, k int) int {
	m := 1 << len(a)
	sum := make([]int, m)
	for i := 1; i < m; i++ {
		p := bits.TrailingZeros(uint(i))
		sum[i] = sum[i&^(1<<p)] + a[p]
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, set int) (res int) {
		if set == 0 {
			return
		}
		if p == k {
			return 1e9
		}
		dv := &dp[p][set]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for sub, ok := set, true; ok; ok = sub != set {
			res = min(res, max(sum[sub], f(p+1, set^sub)))
			sub = (sub - 1) & set
		}
		return
	}
	return f(0, m-1)
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
