package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func interpret(s string) string {
	return strings.NewReplacer("(al)", "al", "()", "o").Replace(s)
}
