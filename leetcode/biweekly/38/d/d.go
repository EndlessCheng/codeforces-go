package main

// github.com/EndlessCheng/codeforces-go
func numWays(words []string, target string) int {
	const mod int = 1e9 + 7
	n, m := len(words[0]), len(target)
	cnt := make([][26]int, n)
	for _, s := range words {
		for i, b := range s {
			cnt[i][b-'a']++
		}
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, j int) (res int) {
		if j == m {
			return 1
		}
		if i == n {
			return 0
		}
		dv := &dp[i][j]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		return (cnt[i][target[j]-'a']*f(i+1, j+1) + f(i+1, j)) % mod
	}
	return f(0, 0) % mod
}
