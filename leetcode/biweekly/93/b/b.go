package main

import (
	"math"
	"sort"
)

// https://space.bilibili.com/206214
func maxStarSum(vals []int, edges [][]int, k int) int {
	g := make([]sort.IntSlice, len(vals))
	for _, e := range edges {
		x, y := e[0], e[1]
		if vals[y] > 0 {
			g[x] = append(g[x], vals[y])
		}
		if vals[x] > 0 {
			g[y] = append(g[y], vals[x])
		}
	}
	ans := math.MinInt32
	for i, a := range g {
		sort.Sort(sort.Reverse(a))
		if len(a) > k {
			a = a[:k]
		}
		sum := vals[i]
		for _, v := range a {
			sum += v
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int { if b > a { return b }; return a }
