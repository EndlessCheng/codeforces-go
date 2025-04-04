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
	return len(nums) - f[3]
}

func minimumOperations2(nums []int) int {
	n := len(nums)
	f := make([][4]int, n+1)
	for i, x := range nums {
		for j := 1; j <= 3; j++ {
			if j < x {
				f[i+1][j] = f[i][j]
			} else {
				f[i+1][j] = max(f[i][j], f[i][x]+1)
			}
		}
	}
	return n - f[n][3]
}

func minimumOperations1(nums []int) int {
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
