package main

import "math/bits"

// https://space.bilibili.com/206214
func countMonobit(n int) int {
	return bits.Len(uint(n + 1))
}
