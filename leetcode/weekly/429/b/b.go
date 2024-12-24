package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func maxDistinctElements(nums []int, k int) (ans int) {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			ans++
			pre = x
		}
	}
	return
}
