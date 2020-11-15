package main

// github.com/EndlessCheng/codeforces-go
func minFallingPathSum(a [][]int) (ans int) {
	n, m := len(a), len(a[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if i == n {
			return
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1e9
		for k, v := range a[i] {
			if k != j {
				res = min(res, v+f(i+1, k))
			}
		}
		return
	}
	return f(0, m)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
