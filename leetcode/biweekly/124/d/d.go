package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func maxSelectedElements(nums []int) (ans int) {
	slices.Sort(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
	}
	for _, res := range f {
		ans = max(ans, res)
	}
	return
}
