package main

import "slices"

// https://space.bilibili.com/206214
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * (j + 1) // 横切
			i++
		} else {
			ans += verticalCut[j] * (i + 1) // 竖切
			j++
		}
	}
	return int64(ans)
}
