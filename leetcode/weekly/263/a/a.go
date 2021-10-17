package main

import (
	"strconv"
	"strings"
	"unicode"
)

// Go 模拟

// github.com/EndlessCheng/codeforces-go
func areNumbersAscending(s string) bool {
	prev := 0
	for _, token := range strings.Split(s, " ") {
		if unicode.IsDigit(rune(token[0])) {
			v, _ := strconv.Atoi(token)
			if v <= prev {
				return false
			}
			prev = v
		}
	}
	return true
}
