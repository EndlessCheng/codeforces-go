package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func arrayStringsAreEqual(word1 []string, word2 []string) (ans bool) {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
