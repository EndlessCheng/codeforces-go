package main

// github.com/EndlessCheng/codeforces-go
func stoneGame(a []int) (ans bool) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(a)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(l, r int) (res int) {
		if l == r {
			return a[l]
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		return max(a[l]-f(l+1, r), a[r]-f(l, r-1))
	}
	return f(0, n-1) > 0
}
