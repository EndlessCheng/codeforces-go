package main

// https://space.bilibili.com/206214
func isCircularSentence(s string) bool {
	if s[0] != s[len(s)-1] {
		return false
	}
	for i, c := range s {
		if c == ' ' && s[i-1] != s[i+1] {
			return false
		}
	}
	return true
}
