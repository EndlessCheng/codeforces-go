package main

// https://space.bilibili.com/206214
func countValidSubstrings(s string) (ans int64) {
	left := -1 // left+1 是有效子串的最小左端点
	for i := range s {
		// s[i-2:i+1] == "abc" 
		if i >= 2 && s[i] == 'c' && s[i-1] == 'b' && s[i-2] == 'a' {
			left = i - 2
		}
		ans += int64(i - left)
	}
	return
}
