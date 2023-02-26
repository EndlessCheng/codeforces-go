package main

import "sort"

// https://space.bilibili.com/206214
func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x {
			i++
		}
	}
	return i * 2
}
