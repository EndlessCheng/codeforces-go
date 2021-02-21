package main

// github.com/EndlessCheng/codeforces-go
func LCS(s, t []byte) [][]int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i, v := range s {
		for j, w := range t {
			if v == w {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp
}

func LPS(s []byte) [][]int {
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
	return dp
}

func longestPalindrome(S, T string) (ans int) {
	s, t := []byte(S), []byte(T)
	n, m := len(s), len(t)
	for i := 0; i < m/2; i++ {
		t[i], t[m-1-i] = t[m-1-i], t[i]
	}

	lcs := LCS(s, t)
	if lcs[n][m] == 0 {
		return
	}
	ans = max(ans, 2*lcs[n][m])

	lps := LPS(t)
	for j, l := range lcs[n][:m] {
		if l > 0 {
			ans = max(ans, 2*l+lps[j][m-1])
		}
	}

	lps = LPS(s)
	for i, row := range lcs[:n] {
		if l := row[m]; l > 0 {
			ans = max(ans, 2*l+lps[i][n-1])
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
