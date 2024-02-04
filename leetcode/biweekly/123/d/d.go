package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func numberOfPairs(points [][]int) (ans int) {
	slices.SortFunc(points, func(p, q []int) int {
		if p[0] != q[0] {
			return p[0] - q[0]
		}
		return q[1] - p[1]
	})
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
