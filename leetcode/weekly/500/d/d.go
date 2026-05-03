package main

import (
	"cmp"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
// 354. 俄罗斯套娃信封问题
func maxEnvelopes(envelopes [][2]int) int {
	slices.SortFunc(envelopes, func(a, b [2]int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	g := []int{}
	for _, e := range envelopes {
		h := e[1]
		j := sort.SearchInts(g, h+1) // 允许 LIS 相邻元素相等
		if j < len(g) {
			g[j] = h
		} else {
			g = append(g, h)
		}
	}
	return len(g)
}

func maxFixedPoints(nums []int) int {
	a := [][2]int{}
	for i, x := range nums {
		if i >= x {
			a = append(a, [2]int{x, i - x})
		}
	}
	return maxEnvelopes(a)
}
