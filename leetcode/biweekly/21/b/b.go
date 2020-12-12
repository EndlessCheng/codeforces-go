package main

import "strings"

// github.com/EndlessCheng/codeforces-go
func findTheLongestSubstring(s string) (ans int) {
	pos, v := map[int]int{0: 0}, 0
	for i, b := range s {
		if strings.Contains("aeiou", string(b)) {
			v ^= 1 << (b - 'a')
		}
		i++
		if p, has := pos[v]; has {
			ans = max(ans, i-p)
		} else {
			pos[v] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}