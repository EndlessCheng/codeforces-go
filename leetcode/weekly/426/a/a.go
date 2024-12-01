package main

import "math/bits"

// https://space.bilibili.com/206214
func smallestNumber(n int) int {
	return 1<<bits.Len(uint(n)) - 1
}
