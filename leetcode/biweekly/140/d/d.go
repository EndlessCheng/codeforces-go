package main

import "slices"

// https://space.bilibili.com/206214
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func rev(s string) string {
	t := []byte(s)
	slices.Reverse(t)
	return string(t)
}

func minStartingIndex(s, pattern string) int {
	preZ := calcZ(pattern + s)
	sufZ := calcZ(rev(pattern) + rev(s))
	slices.Reverse(sufZ) // 也可以不反转，下面写 sufZ[len(sufZ)-i]
	m := len(pattern)
	for i := m; i <= len(s); i++ {
		if preZ[i]+sufZ[i-1] >= m-1 {
			return i - m
		}
	}
	return -1
}
