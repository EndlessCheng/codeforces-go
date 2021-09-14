package main

// github.com/EndlessCheng/codeforces-go
func minDays(n int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := map[int]int{}
	var f func(int) int
	f = func(n int) (res int) {
		if n == 0 {
			return
		}
		if v, ok := dp[n]; ok {
			return v
		}
		defer func() { dp[n] = res }()
		res = 1e9
		if n%6 > 0 { // 6|n 时必然选下面两种决策之一更优
			res = min(res, 1+f(n-1))
		}
		if n%2 == 0 {
			res = min(res, 1+f(n/2))
		}
		if n%3 == 0 {
			res = min(res, 1+f(n/3))
		}
		return
	}
	return f(n)
}
