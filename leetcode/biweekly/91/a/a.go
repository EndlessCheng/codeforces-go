package main

import "sort"

// https://space.bilibili.com/206214
// 点评：排序后就可以一左一右地枚举了，注意不需要除以 2。
func distinctAverages(nums []int) int {
	set := map[int]struct{}{}
	sort.Ints(nums)
	for i, n := 0, len(nums); i < n/2; i++ {
		set[nums[i]+nums[n-1-i]] = struct{}{}
	}
	return len(set)
}
