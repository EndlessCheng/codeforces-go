package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func maxGoodNumber(nums []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int {
		lenA := bits.Len(uint(a))
		lenB := bits.Len(uint(b))
		return (b<<lenA | a) - (a<<lenB | b)
	})

	for _, x := range nums {
		ans = ans<<bits.Len(uint(x)) | x
	}
	return
}
