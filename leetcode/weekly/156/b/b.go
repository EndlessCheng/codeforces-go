package main

// https://space.bilibili.com/206214
func equalSubstring(s, t string, maxCost int) (ans int) {
	l, c := 0, maxCost
	for r := range s {
		c -= abs(int(s[r]) - int(t[r]))
		for c < 0 {
			c += abs(int(s[l]) - int(t[l]))
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
