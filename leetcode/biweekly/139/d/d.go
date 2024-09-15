package main

import "sort"

// https://space.bilibili.com/206214
func maxPathLength(coordinates [][]int, k int) int {
	kx, ky := coordinates[k][0], coordinates[k][1]
	sort.Slice(coordinates, func(i, j int) bool {
		a, b := coordinates[i], coordinates[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

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
	return len(g) + 1
}
