package main

// github.com/EndlessCheng/codeforces-go
func minimumMoves(a []int) (ans int) {
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
		if l >= r {
			return 1
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		if a[l] == a[r] {
			res = f(l+1, r-1)
		}
		for i := l; i < r; i++ {
			res = min(res, f(l, i)+f(i+1, r))
		}
		return
	}
	return f(0, n-1)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
