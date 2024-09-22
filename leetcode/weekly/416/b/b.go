package main

import (
	"math"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxT := slices.Max(workerTimes)
	h := (mountainHeight-1)/len(workerTimes) + 1
	ans := 1 + sort.Search(maxT*h*(h+1)/2, func(m int) bool {
		m++
		leftH := mountainHeight
		for _, t := range workerTimes {
			leftH -= int((math.Sqrt(float64(m/t*8+1)) - 1) / 2)
			if leftH <= 0 {
				return true
			}
		}
		return false
	})
	return int64(ans)
}
