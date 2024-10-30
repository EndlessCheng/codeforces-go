package main

import "math/bits"

// https://space.bilibili.com/206214
func minChanges(n, k int) int {
	if k&^n > 0 {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}
