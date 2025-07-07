package main

import (
	"cmp"
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func numberOfPairs(points [][]int) (ans int) {
	slices.SortFunc(points, func(a, b []int) int { return cmp.Or(a[0]-b[0], b[1]-a[1]) })
	for i, p := range points {
		y0 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y := q[1]
			if y <= y0 && y > maxY {
				maxY = y
				ans++
			}
		}
	}
	return
}
