package main

import "slices"

// https://space.bilibili.com/206214
func maximumTotalSum(maximumHeight []int) int64 {
	slices.SortFunc(maximumHeight, func(a, b int) int { return b - a })
	ans := maximumHeight[0]
	for i := 1; i < len(maximumHeight); i++ {
		maximumHeight[i] = min(maximumHeight[i], maximumHeight[i-1]-1)
		if maximumHeight[i] <= 0 {
			return -1
		}
		ans += maximumHeight[i]
	}
	return int64(ans)
}
