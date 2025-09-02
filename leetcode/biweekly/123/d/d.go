package main

import (
	"cmp"
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func numberOfPairs(points [][]int) (ans int) {
	// x 升序，y 降序
	slices.SortFunc(points, func(a, b []int) int { return cmp.Or(a[0]-b[0], b[1]-a[1]) })
	for i, p := range points {
		y1 := p[1]
		maxY := math.MinInt
		for _, q := range points[i+1:] {
			y2 := q[1]
			if y2 <= y1 && y2 > maxY {
				maxY = y2
				ans++
			}
			if maxY == y1 { // 优化
				break
			}
		}
	}
	return
}
