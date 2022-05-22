package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func percentageLetter(s string, letter byte) int {
	return strings.Count(s, string(letter)) * 100 / len(s)
}
