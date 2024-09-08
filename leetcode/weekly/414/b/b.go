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
	return sort.Search((start[n-1]+d-start[0])/(n-1), func(score int) bool {
		score++
		preX := math.MinInt
		for _, s := range start {
			x := preX + score
			if x > s+d {
				return true
			}
			preX = max(x, s)
		}
		return false
	})
}
