package main

import (
	"cmp"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func sortByReflection(nums []int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		revA := int(bits.Reverse(uint(a)) >> bits.LeadingZeros(uint(a)))
		revB := int(bits.Reverse(uint(b)) >> bits.LeadingZeros(uint(b)))
		return cmp.Or(revA-revB, a-b)
	})
	return nums
}
