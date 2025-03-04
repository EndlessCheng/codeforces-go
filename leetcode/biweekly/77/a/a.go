package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func countPrefixes(words []string, s string) (ans int) {
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			ans++
		}
	}
	return
}
