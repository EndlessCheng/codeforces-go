package main

import "math/bits"

// https://space.bilibili.com/206214
func minOperations(nums []int, k int) int {
	for _, x := range nums {
		k ^= x
	}
	return bits.OnesCount(uint(k))
}
