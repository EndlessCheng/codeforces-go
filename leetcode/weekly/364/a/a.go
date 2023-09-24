package main

import "strings"

// https://space.bilibili.com/206214
func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
