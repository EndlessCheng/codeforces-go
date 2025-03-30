package main

import "math"

// https://space.bilibili.com/206214
func maxActiveSectionsAfterTrade(s string) (ans int) {
	mx := 0
	pre0 := math.MinInt
	cnt := 0
	for i := range len(s) {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			if s[i] == '1' {
				ans += cnt
			} else {
				mx = max(mx, pre0+cnt)
				pre0 = cnt
			}
			cnt = 0
		}
	}
	return ans + mx
}
