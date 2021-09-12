package main

import "strings"

// 模拟

// github.com/EndlessCheng/codeforces-go
func reversePrefix(s string, ch byte) string {
	j := strings.IndexByte(s, ch)
	if j < 0 {
		return s
	}
	t := []byte(s)
	for i := 0; i < j; i++ {
		t[i], t[j] = t[j], t[i]
		j--
	}
	return string(t)
}
