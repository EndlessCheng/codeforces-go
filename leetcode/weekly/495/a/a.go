package main

// https://space.bilibili.com/206214
func firstMatchingIndex(s string) int {
	n := len(s)
	for i := range n/2 + 1 {
		if s[i] == s[n-1-i] {
			return i
		}
	}
	return -1
}
