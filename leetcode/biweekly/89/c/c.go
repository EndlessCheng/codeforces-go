package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimizeArrayValue(nums []int) int {
	return sort.Search(slices.Max(nums), func(limit int) bool {
		extra := 0
		for i := len(nums) - 1; i > 0; i-- {
			extra = max(nums[i]+extra-limit, 0)
		}
		return nums[0]+extra <= limit
	})
}

func minimizeArrayValue2(nums []int) (ans int) {
	s := 0
	for i, x := range nums {
		s += x
		ans = max(ans, (s+i)/(i+1))
	}
	return
}
