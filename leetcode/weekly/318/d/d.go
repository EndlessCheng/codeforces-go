package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	slices.SortFunc(factory, func(a, b []int) int { return a[0] - b[0] })
	slices.Sort(robot)
	m := len(robot)
	f := make([]int, m+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, fa := range factory {
		for j := m; j > 0; j-- {
			for k, cost := 1, 0; k <= min(j, fa[1]); k++ {
				cost += abs(robot[j-k] - fa[0])
				f[j] = min(f[j], f[j-k]+cost)
			}
		}
	}
	return int64(f[m])
}

func abs(x int) int { if x < 0 { return -x }; return x }
