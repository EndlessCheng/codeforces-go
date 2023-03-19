package main

import "sort"

// https://space.bilibili.com/206214
func beautifulSubsets(nums []int, k int) int {
	ans := 1
	groups := map[int]map[int]int{}
	for _, x := range nums {
		if groups[x%k] == nil {
			groups[x%k] = map[int]int{}
		}
		groups[x%k][x]++
	}
	for _, cnt := range groups {
		m := len(cnt)
		type pair struct{ x, c int }
		g := make([]pair, 0, m)
		for x, c := range cnt {
			g = append(g, pair{x, c})
		}
		sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x })
		f := make([]int, m+1)
		f[0] = 1
		f[1] = 1 << g[0].c
		for i := 1; i < m; i++ {
			if g[i].x-g[i-1].x == k {
				f[i+1] = f[i] + f[i-1]*(1<<g[i].c-1)
			} else {
				f[i+1] = f[i] << g[i].c
			}
		}
		ans *= f[m]
	}
	return ans - 1
}
