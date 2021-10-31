package main

import "unicode"

// github.com/EndlessCheng/codeforces-go
func possiblyEquals(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	const mx, bias = 2000, 1000
	dp := make([][][mx]int8, n+1)
	for i := range dp {
		dp[i] = make([][mx]int8, m+1)
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var f func(i, j, d int) int8
	f = func(i, j, d int) (res int8) {
		if i == n && j == m {
			if d == bias { // 匹配成功
				return 1
			}
			return
		}
		dv := &dp[i][j][d]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		// 原始字符串长度相同时，s1[:i] 和 s2[:j] 均可以向后扩展一个字母，要求为相同小写字母
		if d == bias && i < n && j < m && s1[i] == s2[j] && unicode.IsLower(rune(s1[i])) && f(i+1, j+1, bias) > 0 {
			return 1
		}
		if d <= bias && i < n { // s1[:i] 的原始字符串长度不超过 s2[:j] 的原始字符串长度时，扩展 s1[:i]
			if unicode.IsDigit(rune(s1[i])) { // 数字
				for p, v := i, 0; p < n && unicode.IsDigit(rune(s1[p])); p++ {
					v = v*10 + int(s1[p]&15)
					if f(p+1, j, d+v) > 0 {
						return 1
					}
				}
			} else if d < bias && f(i+1, j, d+1) > 0 { // 字符，扩展一位
				return 1
			}
		}
		if d >= bias && j < m { // s2[:j] 的原始字符串长度不超过 s1[:i] 的原始字符串长度时，扩展 s2[:j]
			if unicode.IsDigit(rune(s2[j])) { // 数字
				for q, v := j, 0; q < m && unicode.IsDigit(rune(s2[q])); q++ {
					v = v*10 + int(s2[q]&15)
					if f(i, q+1, d-v) > 0 {
						return 1
					}
				}
			} else if d > bias && f(i, j+1, d-1) > 0 { // 字符，扩展一位
				return 1
			}
		}
		return
	}
	return f(0, 0, bias) > 0
}
