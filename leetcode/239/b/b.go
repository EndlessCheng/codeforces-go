package main

import (
	"strconv"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func splitString(s string) (ans bool) {
o:
	for i := 1; i < len(s); i++ {
		v, _ := strconv.Atoi(s[:i]) // 这个长度溢出也没关系，不影响结果
		v--
		for t := s[i:]; t != ""; v-- {
			for t != "" && t[0] == '0' {
				t = t[1:]
			}
			if t == "" && v == 0 {
				break
			}
			ss := strconv.Itoa(v)
			if !strings.HasPrefix(t, ss) {
				continue o
			}
			t = t[len(ss):]
		}
		return true
	}
	return
}
