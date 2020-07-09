package main

// github.com/EndlessCheng/codeforces-go
func minCost(w1 int, w2 int, packageSum [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int, w1+1)
	for i := range dp {
		dp[i] = make([]int, w2+1)
		for j := range dp[i] {
			dp[i][j] = 1e18
		}
	}
	dp[0][0] = 0
	for _, p := range packageSum {
		p0, p1, cost := p[0], p[1], p[2]
		for i := w1; i >= 0; i-- {
			for j := w2; j >= 0; j-- {
				dp[i][j] = min(dp[i][j], dp[max(i-p0, 0)][max(j-p1, 0)]+cost)
			}
		}
	}
	return dp[w1][w2]
}
