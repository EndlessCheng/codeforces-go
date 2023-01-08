package main

import "sort"

// https://space.bilibili.com/206214
func maximumCount(nums []int) int {
	return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}

func max(a, b int) int { if b > a { return b }; return a }
