package main

import "slices"

// https://space.bilibili.com/206214
func reversePrefix(s string, k int) string {
	t := []byte(s)
	slices.Reverse(t[:k])
	return string(t)
}
