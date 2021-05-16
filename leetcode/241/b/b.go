package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func minSwaps(s string) int {
	ans := int(1e9)
	f := func(tar string) {
		if strings.Count(tar, "0") != strings.Count(s, "0") {
			return
		}
		c := 0
		for i := range s {
			if s[i] == '1' && tar[i] == '0' {
				c++
			}
		}
		if c < ans {
			ans = c
		}
	}
	n := len(s)
	f(strings.Repeat("10", n)[:n])
	f(strings.Repeat("01", n)[:n])
	if ans == 1e9 {
		return -1
	}
	return ans
}
