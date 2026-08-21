package main

import (
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minEatingSpeed(piles []int, h int) int {
	return 1 + sort.Search(slices.Max(piles)-1, func(mid int) bool {
		mid++
		sum := len(piles)
		for _, p := range piles {
			sum += (p - 1) / mid
		}
		return sum <= h
	})
}
