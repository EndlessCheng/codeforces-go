package main

// https://space.bilibili.com/206214
func countKeyChanges(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if s[i-1]&31 != s[i]&31 {
			ans++
		}
	}
	return
}
