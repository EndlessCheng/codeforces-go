package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maximumTastiness(price []int, k int) int {
	slices.Sort(price)
	return sort.Search((price[len(price)-1]-price[0])/(k-1), func(d int) bool {
		d++ // 二分最小的 f(d+1) < k，从而知道最大的 f(d) >= k
		cnt, pre := 1, price[0]
		for _, p := range price[1:] {
			if p >= pre+d {
				cnt++
				pre = p
			}
		}
		return cnt < k
	})
}
