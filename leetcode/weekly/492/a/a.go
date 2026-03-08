package main

import "math"

// https://space.bilibili.com/206214
func minimumIndex(capacity []int, itemSize int) int {
	minC := math.MaxInt
	ans := -1
	for i, c := range capacity {
		if c >= itemSize && c < minC {
			minC = c
			ans = i
		}
	}
	return ans
}
