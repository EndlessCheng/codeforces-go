package main

import "slices"

// https://space.bilibili.com/206214
func minimumArrayLength(nums []int) int {
	m := slices.Min(nums)
	for _, x := range nums {
		if x%m > 0 {
			return 1
		}
	}
	cnt := 0
	for _, x := range nums {
		if x == m {
			cnt++
		}
	}
	return (cnt + 1) / 2
}
