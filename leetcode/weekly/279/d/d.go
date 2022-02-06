package main

// 前后缀分解 + DP

// github.com/EndlessCheng/codeforces-go
func minimumTime(s string) int {
	n := len(s)
	ans := n
	pre := 0
	for i, ch := range s {
		if ch == '1' {
			ans = min(ans, pre+n-i)
			pre = min(pre+2, i+1)
		}
	}
	return min(ans, pre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
