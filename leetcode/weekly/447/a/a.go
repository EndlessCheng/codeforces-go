package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func isInner(a []int, x int) bool {
	i := sort.SearchInts(a, x)
	return 0 < i && i < len(a)-1 // 左右都有建筑
}

func countCoveredBuildings(_ int, buildings [][]int) (ans int) {
	row := map[int][]int{}
	col := map[int][]int{}
	for _, p := range buildings {
		x, y := p[0], p[1]
		row[x] = append(row[x], y)
		col[y] = append(col[y], x)
	}

	for _, a := range row {
		slices.Sort(a)
	}
	for _, a := range col {
		slices.Sort(a)
	}

	for _, p := range buildings {
		x, y := p[0], p[1]
		if isInner(row[x], y) && isInner(col[y], x) {
			ans++
		}
	}
	return
}
