package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxPossibleScore(start []int, d int) int {
	slices.Sort(start)
	n := len(start)
	// 二分最小的不满足要求的 score+1，最终得到的答案就是最大的满足要求的 score
	return sort.Search((start[n-1]+d-start[0])/(n-1), func(score int) bool {
		score++
		x := math.MinInt
		for _, s := range start {
			x = max(x+score, s) // x 必须 >= 区间左端点 s
			if x > s+d {
				return true
			}
		}
		return false
	})
}
