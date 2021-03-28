package main

import (
	"strings"
	"unicode"
)

// github.com/EndlessCheng/codeforces-go
func numDifferentIntegers(word string) int {
	set := map[string]struct{}{}
	for _, s := range strings.FieldsFunc(word, func(r rune) bool { return !unicode.IsDigit(r) }) {
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		set[s] = struct{}{}
	}
	return len(set)
}
