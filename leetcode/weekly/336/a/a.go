package main

import "strings"

// https://space.bilibili.com/206214
func vowelStrings(words []string, left, right int) (ans int) {
	for _, s := range words[left : right+1] {
		if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
			ans++
		}
	}
	return
}
