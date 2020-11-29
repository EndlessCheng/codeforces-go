package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func maxRepeating(ss, s string) (ans int) {
	for i := len(ss) / len(s); ; i-- {
		if strings.Contains(ss, strings.Repeat(s, i)) {
			return i
		}
	}
}
