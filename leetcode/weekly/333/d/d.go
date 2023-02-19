package main

import "bytes"

// https://space.bilibili.com/206214
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	for c := byte('a'); c <= 'z'; c++ {
		i := bytes.IndexByte(s, 0)
		if i < 0 { // 都构造完了
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				s[j] = c
			}
		}
	}
	if bytes.IndexByte(s, 0) >= 0 { // 还有没构造完的
		return ""
	}

	// 直接在原数组上验证
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i == n-1 || j == n-1 {
					if lcp[i][j] != 1 {
						return ""
					}
				} else if lcp[i][j] != lcp[i+1][j+1]+1 {
					return ""
				}
			} else if lcp[i][j] > 0 { // 不相等应该是 0
				return ""
			}
		}
	}
	return string(s)
}
