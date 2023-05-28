package main

import "strings"

// https://space.bilibili.com/206214
func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
