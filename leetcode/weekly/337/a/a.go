package main

import "math/bits"

// https://space.bilibili.com/206214
func evenOddBit(n int) []int {
	const mask = 0x55555555
	return []int{bits.OnesCount(uint(n & mask)), bits.OnesCount(uint(n & (mask << 1)))}
}
