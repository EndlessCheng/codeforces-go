package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
