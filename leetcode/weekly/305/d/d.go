package main

import "slices"

// https://space.bilibili.com/206214
func longestIdealString(s string, k int) int {
	f := [26]int{}
	for _, c := range s {
		c := int(c - 'a')
		f[c] = 1 + slices.Max(f[max(c-k, 0):min(c+k+1, 26)])
	}
	return slices.Max(f[:])
}
