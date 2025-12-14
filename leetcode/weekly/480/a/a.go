package main

import "slices"

// https://space.bilibili.com/206214
func sum(a []int) (s int) {
	for _, x := range a {
		s += x
	}
	return s
}

func absDifference(nums []int, k int) int {
	slices.Sort(nums)
	return sum(nums[len(nums)-k:]) - sum(nums[:k])
}
