package main

import "sort"

// https://space.bilibili.com/206214
func answerQueries(nums, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1] // 前缀和
	}
	for i, q := range queries {
		queries[i] = sort.SearchInts(nums, q+1)
	}
	return queries
}
