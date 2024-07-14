package main

import "slices"

// https://space.bilibili.com/206214
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	cntH, cntV := 1, 1
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * cntH // 横切
			i++
			cntV++ // 需要竖切的蛋糕块增加
		} else {
			ans += verticalCut[j] * cntV // 竖切
			j++
			cntH++ // 需要横切的蛋糕块增加
		}
	}
	return int64(ans)
}
