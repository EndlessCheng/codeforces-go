package main

import "slices"

// https://space.bilibili.com/206214
func maxSum(grid [][]int, limits []int, k int) (ans int64) {
	a := []int{}
	cmp := func(a, b int) int { return b - a }
	for i, row := range grid {
		slices.SortFunc(row, cmp)
		a = append(a, row[:limits[i]]...)
	}
	slices.SortFunc(a, cmp)
	for _, x := range a[:k] {
		ans += int64(x)
	}
	return
}

func maxSum1(grid [][]int, limits []int, k int) (ans int64) {
	type pair struct{ x, i int }
	a := make([]pair, 0, len(grid)*len(grid[0])) // 预分配空间
	for i, row := range grid {
		for _, x := range row {
			a = append(a, pair{x, i})
		}
	}
	slices.SortFunc(a, func(a, b pair) int { return b.x - a.x })

	for _, p := range a {
		if k == 0 {
			break
		}
		if limits[p.i] > 0 {
			limits[p.i]--
			k--
			ans += int64(p.x)
		}
	}
	return
}
