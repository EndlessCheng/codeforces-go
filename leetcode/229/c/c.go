package main

// github.com/EndlessCheng/codeforces-go
func maximumScore(a []int, multipliers []int) (ans int) {
	n, m := len(a), len(multipliers)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(l, r int) (res int) {
		if l+r == m {
			return
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		mul := multipliers[l+r]
		return max(a[l]*mul+f(l+1, r), a[n-1-r]*mul+f(l, r+1))
	}
	return f(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
