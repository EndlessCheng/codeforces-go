package main

import "bytes"

// https://space.bilibili.com/206214
func findTheString(lcp [][]int) string {
	n := len(lcp)
	s := make([]byte, n)
	for c := byte('a'); c <= 'z'; c++ {
		i := bytes.IndexByte(s, 0)
		if i < 0 { // 构造完毕
			break
		}
		for j := i; j < n; j++ {
			if lcp[i][j] > 0 {
				s[j] = c
			}
		}
	}
	if bytes.IndexByte(s, 0) >= 0 { // 没有构造完
		return ""
	}

	// 直接在原数组上验证
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			actualLCP := 0
			if s[i] == s[j] {
				actualLCP = 1
				if i < n-1 && j < n-1 {
					actualLCP += lcp[i+1][j+1]
				}
			}
			if lcp[i][j] != actualLCP {
				return ""
			}
		}
	}
	return string(s)
}
