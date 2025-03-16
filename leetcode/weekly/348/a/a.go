package main

import "math/bits"

// https://space.bilibili.com/206214
func minimizedStringLength(s string) int {
	mask := uint(0)
	for _, c := range s {
		mask |= 1 << (c - 'a')
	}
	return bits.OnesCount(mask)
}

func minimizedStringLength1(s string) int {
	set := map[rune]struct{}{}
	for _, c := range s {
		set[c] = struct{}{}
	}
	return len(set)
}
