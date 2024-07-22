package main

import "strings"

// https://space.bilibili.com/206214
func doesAliceWin(s string) bool {
	return strings.ContainsAny(s, "aeiou")
}
