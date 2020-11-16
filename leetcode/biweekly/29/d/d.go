package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minNumberOfSemesters(n int, dependencies [][]int, k int) (ans int) {
	pre := make([]int, n)
	for _, d := range dependencies {
		pre[d[1]-1] |= 1 << (d[0] - 1)
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(subset int) (res int) {
		if subset == 0 {
			return
		}
		dv := &dp[subset]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
	o:
		for sub := subset; sub > 0; sub = (sub - 1) & subset {
			if bits.OnesCount(uint(sub)) <= k {
				for mask := uint(sub); mask > 0; mask &= mask - 1 {
					if pre[bits.TrailingZeros(mask)]&subset > 0 {
						continue o
					}
				}
				res = min(res, f(subset^sub))
			}
		}
		return res + 1
	}
	return f(1<<n - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
