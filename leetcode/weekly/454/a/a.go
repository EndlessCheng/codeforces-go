package main

import (
	"strings"
	"unicode"
)

// https://space.bilibili.com/206214
func generateTag(caption string) string {
	s := strings.ToLower(caption)
	s = strings.Title(s)
	s = strings.ReplaceAll(s, " ", "")
	if s == "" {
		return "#"
	}
	s = "#" + string(unicode.ToLower(rune(s[0]))) + s[1:]
	if len(s) >= 100 {
		s = s[:100]
	}
	return s
}
