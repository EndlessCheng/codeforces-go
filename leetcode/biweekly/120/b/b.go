package main

import "slices"

// https://space.bilibili.com/206214
func largestPerimeter(nums []int) int64 {
	slices.Sort(nums)
	ans := -1
	s := 0
	for _, x := range nums {
		s += x
		if s > x*2 { // s-x > x
			ans = s
		}
	}
	return int64(ans)
}
