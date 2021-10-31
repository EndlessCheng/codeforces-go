package main

import "unicode"

// github.com/EndlessCheng/codeforces-go
func possiblyEquals(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	const mx, bias = 2000, 1000
	vis := make([][][mx]bool, n+1)
	for i := range vis {
		vis[i] = make([][mx]bool, m+1)
	}
	var dfs func(i, j, d int) bool
	dfs = func(i, j, d int) bool {
		if i == n && j == m {
			return d == 0 // 匹配成功
		}
		if vis[i][j][d+bias] {
			return false
		}
		vis[i][j][d+bias] = true
		// 原始字符串长度相同时，若 s1[i] == s2[j]，则 s1[:i] 和 s2[:j] 均可以向后扩展一个字母
		if d == 0 && i < n && j < m && s1[i] == s2[j] && dfs(i+1, j+1, 0) {
			return true
		}
		if d <= 0 && i < n { // s1[:i] 的原始字符串长度不超过 s2[:j] 的原始字符串长度时，扩展 s1[:i]
			if unicode.IsDigit(rune(s1[i])) { // 数字
				for p, v := i, 0; p < n && unicode.IsDigit(rune(s1[p])); p++ {
					v = v*10 + int(s1[p]&15)
					if dfs(p+1, j, d+v) {
						return true
					}
				}
			} else if d < 0 && dfs(i+1, j, d+1) { // 字符，扩展一位，注意这里需要 d != bias
				return true
			}
		}
		if d >= 0 && j < m { // s2[:j] 的原始字符串长度不超过 s1[:i] 的原始字符串长度时，扩展 s2[:j]
			if unicode.IsDigit(rune(s2[j])) { // 数字
				for q, v := j, 0; q < m && unicode.IsDigit(rune(s2[q])); q++ {
					v = v*10 + int(s2[q]&15)
					if dfs(i, q+1, d-v) {
						return true
					}
				}
			} else if d > 0 && dfs(i, j+1, d-1) { // 字符，扩展一位，注意这里需要 d != bias
				return true
			}
		}
		return false
	}
	return dfs(0, 0, 0)
}
