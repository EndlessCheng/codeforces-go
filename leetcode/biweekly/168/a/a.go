package main

import "slices"

// https://space.bilibili.com/206214
func lexSmallest(s string) string {
	n := len(s)
	ans := s // k = 1 时，操作不改变 s
	for k := 2; k <= n; k++ {
		t := []byte(s[:k])
		slices.Reverse(t)
		ans = min(ans, string(t)+s[k:])

		t = []byte(s[n-k:])
		slices.Reverse(t)
		ans = min(ans, s[:n-k]+string(t))
	}
	return ans
}
