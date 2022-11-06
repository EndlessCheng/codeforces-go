package main

import "sort"

// https://space.bilibili.com/206214
func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Slice(factory, func(i, j int) bool { return factory[i][0] < factory[j][0] })
	sort.Ints(robot)
	m := len(robot)
	f := make([]int, m+1)
	for i := range f {
		f[i] = 1e18
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
