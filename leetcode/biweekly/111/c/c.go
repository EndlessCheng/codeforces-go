package main

import "sort"

// https://space.bilibili.com/206214
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		f[x]++
		f[2] = max(f[2], f[1])
		f[3] = max(f[3], f[2])
	}
	return len(nums) - max(max(f[1], f[2]), f[3])
}

func max(a, b int) int { if b > a { return b }; return a }

func minimumOperationsLIS(nums []int) int {
	g := []int{}
	for _, x := range nums {
		p := sort.SearchInts(g, x+1)
		if p < len(g) {
			g[p] = x
		} else {
			g = append(g, x)
		}
	}
	return len(nums) - len(g)
}
