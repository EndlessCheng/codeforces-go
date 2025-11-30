package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func countElements(nums []int, k int) int {
	n := len(nums)
	if k == 0 {
		return n
	}
	slices.Sort(nums)
	return sort.SearchInts(nums, nums[n-k])
}
