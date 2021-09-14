package main

// github.com/EndlessCheng/codeforces-go
func stoneGameVII(a []int) int {
	n := len(a)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	var f func(int, int) int
	f = func(l, r int) (res int) {
		if l == r {
			return
		}
		dv := &dp[l][r]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		return max(sum[r+1]-sum[l+1]-f(l+1, r), sum[r]-sum[l]-f(l, r-1))
	}
	return f(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
