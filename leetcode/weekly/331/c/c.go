package main

import "sort"

// https://space.bilibili.com/206214
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int { if b > a { return b }; return a }
