package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func findMaxVal(n int, restrictions [][]int, diff []int) int {
	maxVal := make([]int, n)
	for i := range maxVal {
		maxVal[i] = math.MaxInt
	}
	for _, r := range restrictions {
		maxVal[r[0]] = r[1]
	}

	a := make([]int, n)
	for i, d := range diff {
		a[i+1] = min(a[i]+d, maxVal[i+1])
	}
	for i := n - 2; i > 0; i-- {
		a[i] = min(a[i], a[i+1]+diff[i])
	}
	return slices.Max(a)
}
