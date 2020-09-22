package main

// github.com/EndlessCheng/codeforces-go
func maxProductPath(g [][]int) (ans int) {
	const mod int = 1e9 + 7
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
	doMax := func(v [2]int, w int) int { return max(v[0]*w, v[1]*w) }
	doMin := func(v [2]int, w int) int { return min(v[0]*w, v[1]*w) }

	n, m := len(g), len(g[0])
	dp := make([][][2]int, n) // max, min
	for i := range dp {
		dp[i] = make([][2]int, m)
	}
	for i, v := range g[0] {
		if i > 0 {
			v *= dp[0][i-1][0]
		}
		dp[0][i] = [2]int{v, v}
	}
	for i := 1; i < n; i++ {
		for j, v := range g[i] {
			dp[i][j][0] = doMax(dp[i-1][j], v)
			dp[i][j][1] = doMin(dp[i-1][j], v)
			if j > 0 {
				dp[i][j][0] = max(dp[i][j][0], doMax(dp[i][j-1], v))
				dp[i][j][1] = min(dp[i][j][1], doMin(dp[i][j-1], v))
			}
		}
	}

	ans = dp[n-1][m-1][0]
	if ans < 0 {
		return -1
	}
	ans %= mod
	return
}
