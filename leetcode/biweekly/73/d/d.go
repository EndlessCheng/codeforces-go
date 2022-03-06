package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func minMovesToMakePalindrome(s string) (ans int) {
	for s != "" {
		i := strings.IndexByte(s, s[len(s)-1])
		if i == len(s)-1 { // 只出现一次的字符
			ans += i / 2 // 交换到回文中心上
		} else {
			ans += i // 交换到字符串开头
			s = s[:i] + s[i+1:] // 移除 s[i]
		}
		s = s[:len(s)-1] // 移除最后一个字符
	}
	return
}
