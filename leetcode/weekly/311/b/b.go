package main

// https://space.bilibili.com/206214
func longestContinuousSubstring(s string) (ans int) {
	start := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1]+1 {
			ans = max(ans, i-start)
			start = i
		}
	}
	return max(ans, len(s)-start)
}

func max(a, b int) int { if b > a { return b }; return a }
