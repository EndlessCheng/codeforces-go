package main

import "sort"

// https://space.bilibili.com/206214
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		cnt, n := 0, len(nums)
		for i := 0; i < n; i++ {
			if nums[i] <= mx {
				cnt++ // 偷 nums[i]
				i++   // 跳过下一间房子
			}
		}
		return cnt >= k
	})
}
