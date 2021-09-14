package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func halvesAreAlike(s string) (ans bool) {
	s = strings.ToLower(s)
	cnt := func(s string) (c int) {
		for _, b := range s {
			if strings.Contains("aeiou", string(b)) {
				c++
			}
		}
		return
	}
	return cnt(s[:len(s)/2]) == cnt(s[len(s)/2:])
}
