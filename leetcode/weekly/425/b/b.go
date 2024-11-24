package main

import "slices"

// https://space.bilibili.com/206214
func isPossibleToRearrange(s, t string, k int) bool {
	a := make([]string, 0, k) // 预分配空间
	b := make([]string, 0, k)
	n := len(s)
	k = n / k
	for i := k; i <= n; i += k {
		a = append(a, s[i-k:i])
		b = append(b, t[i-k:i])
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
