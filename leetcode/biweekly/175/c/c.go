package main

import (
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func longestSubsequence(nums []int) (ans int) {
	w := bits.Len(uint(slices.Max(nums)))
	for i := range w {
		// 300. 最长递增子序列
		f := []int{}
		for _, x := range nums {
			if x>>i&1 == 0 {
				continue
			}
			j := sort.SearchInts(f, x)
			if j < len(f) {
				f[j] = x
			} else {
				f = append(f, x)
			}
		}
		ans = max(ans, len(f))
	}
	return
}
