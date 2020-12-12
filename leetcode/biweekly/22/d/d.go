package main

// github.com/EndlessCheng/codeforces-go
func maxSizeSlices(a []int) (ans int) {
	n := len(a)
	m := n / 3
	f := func(a []int) int {
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, m+1)
		}
		for i := 1; i < n; i++ {
			for j := 1; j <= m; j++ {
				v := 0
				if i > 1 {
					v = dp[i-2][j-1]
				}
				dp[i][j] = max(dp[i-1][j], v+a[i-1])
			}
		}
		return dp[n-1][m]
	}
	return max(f(a[:n-1]), f(a[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
