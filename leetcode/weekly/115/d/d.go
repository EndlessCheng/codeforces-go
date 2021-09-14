package main

// github.com/EndlessCheng/codeforces-go
func minDeletionSize(a []string) int {
	n := len(a[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, pre int) (res int) {
		if p == n {
			return
		}
		if pre == -1 {
			return min(1+f(p+1, pre), f(p+1, p))
		}
		dv := &dp[p][pre]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		res = 1 + f(p+1, pre)
		for _, s := range a {
			if s[p] < s[pre] {
				return
			}
		}
		return min(res, f(p+1, p))
	}
	return f(0, -1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
