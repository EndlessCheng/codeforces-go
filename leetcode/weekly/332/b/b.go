package main

import "sort"

// https://space.bilibili.com/206214
func countFairPairs(nums []int, lower, upper int) (ans int64) {
	sort.Ints(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:j], lower-x)
		ans += int64(r - l)
	}
	return
}
