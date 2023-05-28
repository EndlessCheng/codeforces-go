package main

// https://space.bilibili.com/206214
func minimumCost(s string) (ans int64) {
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i-1] != s[i] {
			ans += int64(min(i, n-i))
		}
	}
	return
}

func min(a, b int) int { if b < a { return b }; return a }
