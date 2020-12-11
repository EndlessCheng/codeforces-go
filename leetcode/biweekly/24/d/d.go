package main

// github.com/EndlessCheng/codeforces-go
func numberOfArrays(s string, k int) (ans int) {
	n := len(s)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(p int) (res int) {
		if p == n {
			return 1
		}
		if s[p] == '0' {
			return
		}
		dv := &dp[p]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for v := 0; p < n; p++ {
			v = v*10 + int(s[p]&15)
			if v > k {
				break
			}
			res += f(p + 1)
		}
		return res % (1e9 + 7)
	}
	return f(0)
}
