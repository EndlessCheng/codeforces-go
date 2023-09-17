package main

import "sort"

// https://space.bilibili.com/206214
func countWays(nums []int) (ans int) {
	sort.Ints(nums)
	if nums[0] > 0 {
		ans++
	}
	for i, x := range nums {
		if x < i+1 && (i == len(nums)-1 || i+1 < nums[i+1]) {
			ans++
		}
	}
	return
}
