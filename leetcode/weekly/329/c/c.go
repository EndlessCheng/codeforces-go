package main

import "strings"

// https://space.bilibili.com/206214
func makeStringsEqual(s, target string) bool {
	return strings.Contains(s, "1") == strings.Contains(target, "1")
}
