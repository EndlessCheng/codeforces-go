package main

import "unicode"

// https://space.bilibili.com/206214
func passwordStrength(password string) (ans int) {
	vis := [128]bool{}
	for _, ch := range password {
		if vis[ch] {
			continue
		}
		vis[ch] = true
		if unicode.IsLower(ch) {
			ans++
		} else if unicode.IsUpper(ch) {
			ans += 2
		} else if unicode.IsDigit(ch) {
			ans += 3
		} else {
			ans += 5
		}
	}
	return
}
