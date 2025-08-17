package main

import "slices"

// https://space.bilibili.com/206214
func perfectPairs1(nums []int) (ans int64) {
	slices.SortFunc(nums, func(a, b int) int { return abs(a) - abs(b) })
	left := 0
	for i, b := range nums {
		for abs(nums[left])*2 < abs(b) {
			left++
		}
		ans += int64(i - left)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func perfectPairs(nums []int) (ans int64) {
	for i, x := range nums {
		if x < 0 {
			nums[i] *= -1
		}
	}

	slices.Sort(nums)
	left := 0
	for j, b := range nums {
		for nums[left]*2 < b {
			left++
		}
		// a=nums[i]，其中 i 最小是 left，最大是 j-1，一共有 j-left 个
		ans += int64(j - left)
	}
	return
}
