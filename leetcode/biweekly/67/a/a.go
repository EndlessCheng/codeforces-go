package main

import "sort"

// 两次排序

// github.com/EndlessCheng/codeforces-go
func maxSubsequence(nums []int, k int) []int {
	id := make([]int, len(nums))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] > nums[id[j]] }) // 按元素值从大到小排序
	sort.Ints(id[:k]) // 对下标从小到大排序
	ans := make([]int, k)
	for i, j := range id[:k] {
		ans[i] = nums[j]
	}
	return ans
}
