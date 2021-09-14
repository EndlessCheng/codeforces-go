package main

import "strings"

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	const mod int = 1e9 + 7
	calcMaxMatchLengths := func(s string) []int {
		n := len(s)
		maxMatchLengths := make([]int, n)
		cnt := 0
		for i := 1; i < n; i++ {
			b := s[i]
			for cnt > 0 && s[cnt] != b {
				cnt = maxMatchLengths[cnt-1]
			}
			if s[cnt] == b {
				cnt++
			}
			maxMatchLengths[i] = cnt
		}
		return maxMatchLengths
	}
	maxMatchLengths := calcMaxMatchLengths(evil)

	// <=s 的好字符串的数目
	calc := func(s string) int {
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, len(evil))
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		var f func(p, match int, isUpper bool) int
		f = func(p, match int, isUpper bool) (cnt int) {
			if match >= len(evil) {
				return
			}
			if p >= n {
				return 1
			}
			dv := &dp[p][match]
			if !isUpper && *dv >= 0 {
				return *dv
			}
			defer func() {
				if !isUpper {
					*dv = cnt
				}
			}()
			up := byte('z')
			if isUpper {
				up = s[p]
			}
			for c := byte('a'); c <= up; c++ {
				m := match
				for m > 0 && evil[m] != c {
					m = maxMatchLengths[m-1]
				}
				if evil[m] == c {
					m++
				}
				cnt = (cnt + f(p+1, m, isUpper && c == up)) % mod
			}
			return
		}
		return f(0, 0, true)
	}
	ans := calc(s2) - calc(s1)
	if !strings.Contains(s1, evil) {
		ans++
	}
	return (ans%mod + mod) % mod
}
