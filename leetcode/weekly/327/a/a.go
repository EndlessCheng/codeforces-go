package main

import "sort"

// https://space.bilibili.com/206214
func maximumCount(nums []int) int {
	return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}
