package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func truncateSentence(s string, k int) string {
	return strings.Join(strings.Split(s, " ")[:k], " ")
}
