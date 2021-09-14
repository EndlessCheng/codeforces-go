package main

// github.com/EndlessCheng/codeforces-go
func profitableSchemes(g int, p int, group []int, profit []int) (ans int) {
	const mod int = 1e9 + 7
	dp := make([][]int, g+1)
	for i := range dp {
		dp[i] = make([]int, p+1)
		dp[i][0] = 1
	}
	for i, w := range group {
		w2 := profit[i]
		for j := g; j >= w; j-- { // g 有上限
			for k := p; k >= 0; k-- { // p 有下限
				dp[j][k] = (dp[j][k] + dp[j-w][max(k-w2, 0)]) % mod
			}
		}
	}
	return dp[g][p]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
