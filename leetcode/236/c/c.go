package main

// github.com/EndlessCheng/codeforces-go
func minSideJumps(a []int) int {
	n := len(a)
	dp := make([][4]int, n)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = 1e9
		}
	}
	dp[0] = [4]int{0, 1, 0, 1}
	for i := 1; i < n; i++ {
		pre, cur := a[i-1], a[i]
		for j := 1; j < 4; j++ {
			if j == cur {
				continue
			}
			if j != pre {
				dp[i][j] = min(dp[i][j], dp[i-1][j])
			}
			for k := 1; k < 4; k++ {
				if k != j && k != cur && k != pre {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+1)
				}
			}
		}
	}
	return min(min(dp[n-1][1], dp[n-1][2]), dp[n-1][3])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
