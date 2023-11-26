package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums[i] - nums[j] })

	ans := make([]int, n)
	for i := 0; i < n; {
		st := i
		for i++; i < n && nums[ids[i]]-nums[ids[i-1]] <= limit; i++ {
		}
		subIds := slices.Clone(ids[st:i])
		slices.Sort(subIds)
		for j, idx := range subIds {
			ans[idx] = nums[ids[st+j]]
		}
	}
	return ans
}
