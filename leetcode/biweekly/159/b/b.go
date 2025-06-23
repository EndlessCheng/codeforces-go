package main

import "math"

// https://space.bilibili.com/206214
func maxArea(coords [][]int) int64 {
	calc := func(j int) (res int) {
		minX, maxX := math.MaxInt, 0
		minY := map[int]int{}
		maxY := map[int]int{}
		for _, p := range coords {
			x, y := p[j], p[1-j]
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
			res = max(res, (maxY[x]-y)*max(maxX-x, x-minX))
		}
		return
	}

	ans := max(calc(0), calc(1))
	if ans == 0 {
		ans = -1
	}
	return int64(ans)
}
