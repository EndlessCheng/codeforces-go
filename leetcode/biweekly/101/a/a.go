package main

import (
	"math/bits"
)

// https://space.bilibili.com/206214
func minNumber(nums1, nums2 []int) int {
	var mask1, mask2 uint
	for _, x := range nums1 { mask1 |= 1 << x }
	for _, x := range nums2 { mask2 |= 1 << x }
	if m := mask1 & mask2; m > 0 {
		return bits.TrailingZeros(m)
	}
	x, y := bits.TrailingZeros(mask1), bits.TrailingZeros(mask2)
	return min(x*10+y, y*10+x)
}

func min(a, b int) int { if b < a { return b }; return a }
