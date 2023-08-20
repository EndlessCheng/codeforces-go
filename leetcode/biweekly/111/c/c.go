package main

import "sort"

// https://space.bilibili.com/206214
func minimumOperations(nums []int) int {
	f := [4]int{}
	for _, x := range nums {
		for j := 3; j > 0; j-- {
			for k := 1; k <= j; k++ {
				f[j] = min(f[j], f[k])
			}
			if j != x {
				f[j]++
			}
		}
	}
	ans := len(nums)
	for _, v := range f[1:] {
		ans = min(ans, v)
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }

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
