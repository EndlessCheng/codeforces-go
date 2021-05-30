package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func minimumXORSum(x, y []int) int {
	m := 1 << len(x)
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 2e9
	}
	dp[0] = 0
	for s, dv := range dp[:m-1] {
		v := x[bits.OnesCount(uint(s))]
		for t, lb := s^(m-1), 0; t > 0; t ^= lb {
			lb = t & -t
			w := y[bits.TrailingZeros(uint(lb))]
			dp[s|lb] = min(dp[s|lb], v^w+dv)
		}
	}
	return dp[m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
