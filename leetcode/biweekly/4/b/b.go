package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func removeVowels(s string) string {
	return strings.NewReplacer("a", "", "e", "", "i", "", "o", "", "u", "").Replace(s)
}
