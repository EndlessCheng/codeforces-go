package main

// github.com/EndlessCheng/codeforces-go
func longestPalindromeSubsequence(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func isValidPalindrome(s string, k int) bool {
	return longestPalindromeSubsequence(s)+k >= len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
