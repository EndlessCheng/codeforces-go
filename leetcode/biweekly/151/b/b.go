package main

import "math"

// https://space.bilibili.com/206214
func countArrays(original []int, bounds [][]int) int {
	mn, mx := math.MinInt, math.MaxInt
	for i, b := range bounds {
		mn = max(mn, b[0]-original[i]) // 计算区间交集
		mx = min(mx, b[1]-original[i])
	}
	return max(mx-mn+1, 0) // 注意交集可能是空的
}
