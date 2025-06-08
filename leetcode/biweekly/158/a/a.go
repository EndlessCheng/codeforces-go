package main

import (
	"maps"
	"slices"
)

// https://space.bilibili.com/206214
func maxSumDistinctTriplet(x, y []int) int {
	mx := map[int]int{}
	for i, v := range x {
		mx[v] = max(mx[v], y[i])
	}
	if len(mx) < 3 {
		return -1
	}
	a := slices.SortedFunc(maps.Values(mx), func(a, b int) int { return b - a })
	return a[0] + a[1] + a[2]
}
