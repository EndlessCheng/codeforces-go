package main

import "strconv"

// github.com/EndlessCheng/codeforces-go
func digitsCount(D, lower, upper int) (ans int) {
	tar := byte('0' + D)
	calc := func(S int) int {
		if S == 0 {
			return 0
		}
		const lowerC, upperC byte = '0', '9'
		s := strconv.Itoa(S)
		n := len(s)
		const mxL = 11
		dp := make([][mxL][2]int, n)
		for i := range dp {
			for j := 0; j < mxL; j++ {
				dp[i][j] = [2]int{-1, -1}
			}
		}
		var f func(p, sum, hasD int, limitUp bool) int
		f = func(p, sum, hasD int, limitUp bool) (res int) {
			if p == n {
				return sum
			}
			if !limitUp {
				dv := &dp[p][sum][hasD]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := upperC
			if limitUp {
				up = s[p]
			}
			lw := lowerC
			if hasD == 0 {
				res += f(p+1, 0, 0, limitUp && '0' == up)
				lw++
			}
			for d := lw; d <= up; d++ {
				tmp := sum
				if d == tar {
					tmp++
				}
				res += f(p+1, tmp, 1, limitUp && d == up)
			}
			return
		}
		return f(0, 0, 0, true)
	}
	ansLower := calc(lower - 1)
	ansUpper := calc(upper)
	ans = ansUpper - ansLower
	return
}
