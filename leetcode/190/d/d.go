package main

func maxDotProduct(a, b []int) (ans int) {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1e18
		}
	}
	for i, v := range a {
		for j, w := range b {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1], dp[i][j]+v*w, v*w)
		}
	}
	return dp[n][m]
}
