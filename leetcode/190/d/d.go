package main

func maxDotProduct(s1, s2 []int) (ans int) {
	max := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = -1e18
		}
	}
	for i, v := range s1 {
		for j, w := range s2 {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1], dp[i][j]+v*w, v*w)
		}
	}
	return dp[n][m]
}
