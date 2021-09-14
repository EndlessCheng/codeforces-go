package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func isPrefixString(s string, words []string) bool {
	sum := 0
	for i, word := range words {
		sum += len(word)
		if sum < len(s) {
			continue
		}
		if sum > len(s) {
			return false
		}
		return strings.Join(words[:i+1], "") == s
	}
	return false
}
