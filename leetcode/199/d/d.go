package main

var w = [...]int{1: 1, 100: 4}

func init() {
	for i := 2; i < 10; i++ {
		w[i] = 2
	}
	for i := 10; i < 100; i++ {
		w[i] = 3
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// github.com/EndlessCheng/codeforces-go
func getLengthOfOptimalCompression(s string, K int) (ans int) {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, K+1)
		for j := range dp[i] {
			dp[i][j] = 1e9
		}
	}
	dp[0][0] = 0
	for i := range s {
		b := s[i]
		for k := 0; k <= K; k++ {
			if k > 0 {
				dp[i+1][k] = min(dp[i+1][k], dp[i][k-1])
			}
			diff, same := 0, 0
			for j := i; j >= 0; j-- {
				if s[j] != b {
					if diff++; diff > k {
						break
					}
				} else {
					same++
				}
				dp[i+1][k] = min(dp[i+1][k], dp[j][k-diff]+w[same])
			}
		}
	}
	return dp[n][K]
}
