package main

import "slices"

// https://space.bilibili.com/206214
func minimizeSum(a []int) int {
	slices.Sort(a)
	n := len(a)
	return min(a[n-3]-a[0], a[n-2]-a[1], a[n-1]-a[2])
}
