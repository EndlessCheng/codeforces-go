package main

import "sort"

// https://space.bilibili.com/206214
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	return sort.Search(1e9, func(mx int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= mx {
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}
