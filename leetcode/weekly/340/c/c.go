package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minimizeMax(nums []int, p int) int {
	slices.Sort(nums)

	n := len(nums)
	diffs := make([]int, n)
	for i := 1; i < n; i++ {
		diffs[i] = nums[i] - nums[i-1]
	}
	slices.Sort(diffs)

	idx := sort.Search(n-1, func(idx int) bool {
		mx := diffs[idx]
		cnt := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] <= mx { // 都选
				cnt++
				i++
			}
		}
		return cnt >= p
	})
	return diffs[idx]
}
