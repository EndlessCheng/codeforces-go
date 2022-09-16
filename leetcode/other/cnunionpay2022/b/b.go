package main

import "sort"

// https://space.bilibili.com/206214
func explorationSupply(a []int, pos []int) []int {
	ans := make([]int, len(pos))
	for i, p := range pos {
		j := sort.SearchInts(a, p)
		if j == len(a) || j > 0 && p-a[j-1] <= a[j]-p {
			j--
		}
		ans[i] = j
	}
	return ans
}
