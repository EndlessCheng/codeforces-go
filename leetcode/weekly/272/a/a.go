package main

// 遍历，找第一个回文串

// github.com/EndlessCheng/codeforces-go
func firstPalindrome(words []string) string {
next:
	for _, w := range words {
		for i, n := 0, len(w); i < n/2; i++ {
			if w[i] != w[n-1-i] {
				continue next
			}
		}
		return w
	}
	return ""
}
