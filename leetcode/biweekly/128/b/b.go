package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int { return p[0] - q[0] })
	x2 := -1
	for _, p := range points {
		if p[0] > x2 {
			ans++
			x2 = p[0] + w
		}
	}
	return
}
