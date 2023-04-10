package main

import "sort"

// https://space.bilibili.com/206214
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	n := len(nums)
	return sort.Search(nums[n-1]-nums[0], func(mx int) bool {
		cnt := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] <= mx { // 都选
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
