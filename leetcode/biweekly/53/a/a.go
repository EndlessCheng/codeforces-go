package main

// github.com/EndlessCheng/codeforces-go
func countGoodSubstrings(s string) (ans int) {
	for i := 2; i < len(s); i++ {
		if s[i-2] != s[i-1] && s[i-2] != s[i] && s[i-1] != s[i] {
			ans++
		}
	}
	return
}
