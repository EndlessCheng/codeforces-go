package main

import "sort"

// https://space.bilibili.com/206214/dynamic
func maximumImportance(n int, roads [][]int) (ans int64) {
	deg := make([]int, n)
	for _, r := range roads {
		deg[r[0]]++
		deg[r[1]]++
	}
	sort.Ints(deg)
	for i, d := range deg {
		ans += int64(i+1) * int64(d)
	}
	return
}
