package main

import "slices"

// https://space.bilibili.com/206214
func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)
	s := 0
	for _, x := range nums {
		s += x
	}
	for i := len(nums) - 1; i > 1; i-- {
		x := nums[i]
		if s > x*2 { // s-x > x
			return int64(s)
		}
		s -= x
	}
	return -1
}
