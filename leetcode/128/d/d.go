package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func numDupDigitsAtMostN(N int) (ans int) {
	s := strconv.Itoa(N)
	n := len(s)
	dp := make([][1024]int, n)
	for i := range dp {
		for j := 0; j < 1024; j++ {
			dp[i][j] = -1
		}
	}
	var f func(p, set int, isUpper bool) int
	f = func(p, set int, isUpper bool) (res int) {
		if p == n {
			return 1
		}
		if !isUpper {
			dv := &dp[p][set]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		up := 9
		if isUpper {
			up = int(s[p] & 15)
		}
		for d := 0; d <= up; d++ {
			if d == 0 && set == 0 {
				res += f(p+1, 0, isUpper && d == up)
			} else if set>>d&1 == 0 {
				res += f(p+1, set|1<<d, isUpper && d == up)
			}
		}
		return
	}
	return N + 1 - f(0, 0, true)
}
