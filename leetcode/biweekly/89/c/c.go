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
			newNum := nums[i] + extra    // 把右边的积木堆到 nums[i] 上
			extra = max(newNum-limit, 0) // 如果 newNum-limit > 0，那么多出的积木继续丢给左边
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
