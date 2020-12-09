package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func atMostNGivenDigitSet(ds []string, N int) (ans int) {
	s := strconv.Itoa(N)
	n := len(s)
	dp := make([][2]int, n)
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	var f func(p, hasD int, isUpper bool) int
	f = func(p, hasD int, isUpper bool) (res int) {
		if p == n {
			return hasD
		}
		if !isUpper {
			dv := &dp[p][hasD]
			if *dv >= 0 {
				return *dv
			}
			defer func() { *dv = res }()
		}
		up := byte('9')
		if isUpper {
			up = s[p]
		}
		if hasD == 0 {
			res += f(p+1, 0, isUpper && '0' == up)
		}
		for _, d := range ds {
			if d[0] > up {
				break
			}
			res += f(p+1, 1, isUpper && d[0] == up)
		}
		return
	}
	return f(0, 0, true)
}
