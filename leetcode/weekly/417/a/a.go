package main

import "math/bits"

// https://space.bilibili.com/206214
func kthCharacter(k int) byte {
	return 'a' + byte(bits.OnesCount(uint(k-1)))
}
