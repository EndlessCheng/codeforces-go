package main

import "sort"

// https://space.bilibili.com/206214
func countPairs(nums []int, target int) (ans int) {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}
