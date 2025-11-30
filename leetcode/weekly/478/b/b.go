package main

import "math/bits"

// https://space.bilibili.com/206214
func maxDistinct(s string) int {
	vis := 0
	for _, c := range s {
		vis |= 1 << (c - 'a')
	}
	return bits.OnesCount(uint(vis))
}
