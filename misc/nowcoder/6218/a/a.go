package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func change(s string) string {
	return strings.ReplaceAll(s, "a", "") + strings.Repeat("a", strings.Count(s, "a"))
}
