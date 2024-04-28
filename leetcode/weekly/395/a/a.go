package main

import "slices"

// https://space.bilibili.com/206214
func addedInteger(nums1, nums2 []int) int {
	return slices.Min(nums2) - slices.Min(nums1)
}
