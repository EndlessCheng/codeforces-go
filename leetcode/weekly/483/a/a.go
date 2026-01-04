package main

import "strings"

// https://space.bilibili.com/206214
func largestEven(s string) string {
	return strings.TrimRight(s, "1")
}
