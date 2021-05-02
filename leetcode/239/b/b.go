package main

import (
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func splitString(s string) bool {
o:
	for i := 1; i < len(s); i++ {
		v, _ := strconv.Atoi(s[:i]) // 这个长度溢出也没关系，不影响结果
		v--
		for t := s[i:]; t != ""; v-- {
			// 移除前导零，至少保留一个数字
			for len(t) > 1 && t[0] == '0' {
				t = t[1:]
			}
			// 判断 v 是否符合
			sv := strconv.Itoa(v)
			if !strings.HasPrefix(t, sv) {
				continue o
			}
			t = t[len(sv):]
		}
		return true
	}
	return false
}
