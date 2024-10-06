package main

import (
	"cmp"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxPathLength(coordinates [][]int, k int) int {
	kx, ky := coordinates[k][0], coordinates[k][1]
	slices.SortFunc(coordinates, func(a, b []int) int { return cmp.Or(a[0]-b[0], b[1]-a[1]) })

	g := []int{}
	for _, p := range coordinates {
		x, y := p[0], p[1]
		if x < kx && y < ky || x > kx && y > ky {
			j := sort.SearchInts(g, y)
			if j < len(g) {
				g[j] = y
			} else {
				g = append(g, y)
			}
		}
	}
	return len(g) + 1 // 算上 coordinates[k]
}
