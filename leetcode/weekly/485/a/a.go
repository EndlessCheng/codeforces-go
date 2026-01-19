package main

import (
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
func vowelConsonantScore(s string) int {
	v, c := 0, 0
	for _, ch := range s {
		if !unicode.IsLetter(ch) {
			continue
		}
		if strings.ContainsRune("aeiou", ch) {
			v++
		} else {
			c++
		}
	}

	if c > 0 {
		return v / c
	}
	return 0
}
