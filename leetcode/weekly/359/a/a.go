package main

// https://space.bilibili.com/206214
func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}
