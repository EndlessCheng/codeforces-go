package main

// https://space.bilibili.com/206214
func countValidSubstrings(s string) (ans int64) {
	left := -1 // left+1 是有效子串的最小左端点
	for i := range s {
		if i >= 2 && s[i-2:i+1] == "abc" {
			left = i - 2
		}
		ans += int64(i - left)
	}
	return
}
