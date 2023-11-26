package main

import "strings"

// https://space.bilibili.com/206214
func findWordsContaining(words []string, x byte) (ans []int) {
	for i, s := range words {
		if strings.IndexByte(s, x) >= 0 {
			ans = append(ans, i)
		}
	}
	return
}
