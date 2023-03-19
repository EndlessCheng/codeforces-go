package main

import "math/bits"

// https://space.bilibili.com/206214
func evenOddBit(n int) []int {
	const mask = 0x5555
	return []int{bits.OnesCount16(uint16(n & mask)), bits.OnesCount16(uint16(n & (mask >> 1)))}
}
