package main

// https://space.bilibili.com/206214
func findTheLongestBalancedSubstring(s string) (ans int) {
	pre, cur := 0, 0
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] {
			if c == '1' {
				ans = max(ans, min(pre, cur)*2)
			}
			pre = cur
			cur = 0
		}
	}
	return
}

func min(a, b int) int { if a > b { return b }; return a }
func max(a, b int) int { if a < b { return b }; return a }
