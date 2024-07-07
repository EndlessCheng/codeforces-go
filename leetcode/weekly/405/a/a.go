package main

// https://space.bilibili.com/206214
func getEncryptedString(s string, k int) string {
	k %= len(s)
	return s[k:] + s[:k]
}
