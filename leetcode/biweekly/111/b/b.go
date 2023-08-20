package main

// https://space.bilibili.com/206214
func canMakeSubsequence(s, t string) bool {
	if len(s) < len(t) {
		return false
	}
	j := 0
	for _, b := range s {
		c := byte(b) + 1
		if b == 'z' {
			c = 'a'
		}
		if byte(b) == t[j] || c == t[j] {
			j++
			if j == len(t) {
				return true
			}
		}
	}
	return false
}
