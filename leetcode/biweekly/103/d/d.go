package main

import "sort"

// https://space.bilibili.com/206214
func countOperationsToEmptyArray(nums []int) int64 {
	n := len(nums)
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] < nums[id[j]] })

	ans := n // 先把 n 计入答案
	for k := 1; k < n; k++ {
		if id[k] < id[k-1] {
			ans += n - k
		}
	}
	return int64(ans)
}
