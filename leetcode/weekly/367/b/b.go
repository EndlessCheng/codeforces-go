package main

import (
	"strings"
)

// https://space.bilibili.com/206214
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt1 := 0
	left := 0
	for right, b := range s {
		cnt1 += int(b & 1)
		for cnt1 > k || s[left] == '0' {
			cnt1 -= int(s[left] & 1)
			left++
		}
		if cnt1 == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}