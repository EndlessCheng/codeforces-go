package main

import (
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
func largestPower(nums []int) []int {
	ans := [15]int{}
	slices.SortFunc(nums, func(a, b int) int { return b - a })
	maxWidth := bits.Len(uint(nums[0]))
	for i := maxWidth - 1; i >= 0; i-- {
		// 找最长前缀连续 1
		j := 0
		for j < len(nums) && nums[j]>>i&1 > 0 {
			j++
		}
		ans[14-i] = j

		// [0, j-1] 这一位都是 1，后面无关紧要，为方便排序，全置为 0
		for ; j < len(nums); j++ {
			nums[j] &^= 1 << i
		}
		slices.SortFunc(nums, func(a, b int) int { return b - a })
	}
	return ans[:]
}
