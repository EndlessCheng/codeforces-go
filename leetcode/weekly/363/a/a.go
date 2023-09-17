package main

import "math/bits"

// https://space.bilibili.com/206214
func sumIndicesWithKSetBits(nums []int, k int) (ans int) {
	for i, x := range nums {
		if bits.OnesCount(uint(i)) == k {
			ans += x
		}
	}
	return
}
