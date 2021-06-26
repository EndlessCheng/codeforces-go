package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func removeOccurrences(s, part string) string {
	for {
		i := strings.Index(s, part)
		if i < 0 {
			return s
		}
		s = s[:i] + s[i+len(part):]
	}
}
