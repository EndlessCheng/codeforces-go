package main

import "math"

// https://space.bilibili.com/206214
func maxArea(coords [][]int) int64 {
	ans := 0

	calc := func() {
		minX, maxX := math.MaxInt, 0
		minY := map[int]int{}
		maxY := map[int]int{}
		for _, p := range coords {
			x, y := p[0], p[1]
			minX = min(minX, x)
			maxX = max(maxX, x)
			maxY[x] = max(maxY[x], y)
			mn, ok := minY[x]
			if !ok {
				minY[x] = y
			} else {
				minY[x] = min(mn, y)
			}
		}
		for x, y := range minY {
			ans = max(ans, (maxY[x]-y)*max(maxX-x, x-minX))
		}
	}
	calc()

	for _, p := range coords {
		p[0], p[1] = p[1], p[0]
	}
	calc()

	if ans == 0 {
		ans = -1
	}
	return int64(ans)
}
