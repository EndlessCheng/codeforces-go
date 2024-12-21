package main

import "slices"

// https://space.bilibili.com/206214
func minimumCost(m, n int, horizontalCut, verticalCut []int) (ans int64) {
	slices.Sort(horizontalCut)
	slices.Sort(verticalCut)
	i, j := 0, 0
	for range m + n - 2 {
		if j == n-1 || i < m-1 && horizontalCut[i] < verticalCut[j] {
			ans += int64(horizontalCut[i] * (n - j)) // 上下连边
			i++
		} else {
			ans += int64(verticalCut[j] * (m - i)) // 左右连边
			j++
		}
	}
	return
}
