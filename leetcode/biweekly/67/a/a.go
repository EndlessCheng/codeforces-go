package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maxSubsequence(nums []int, k int) []int {
	// 创建下标数组，对下标数组排序
	idx := make([]int, len(nums))
	for i := range idx {
		idx[i] = i
	}
	slices.SortFunc(idx, func(i, j int) int { return nums[j] - nums[i] })

	// 取前 k 大元素的下标，排序
	idx = idx[:k]
	slices.Sort(idx)

	// 取出 nums 的子序列
	for i, j := range idx {
		idx[i] = nums[j]
	}
	return idx
}
