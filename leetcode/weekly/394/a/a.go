package main

import "math/bits"

// https://space.bilibili.com/206214
func numberOfSpecialChars(word string) int {
	mask := [2]int{}
	for _, c := range word {
		mask[c>>5&1] |= 1 << (c & 31)
	}
	return bits.OnesCount(uint(mask[0] & mask[1]))
}
