package main

import (
	"math"
	"slices"
)

// github.com/EndlessCheng/codeforces-go
func minimumAbsDifference(arr []int) (ans [][]int) {
	slices.Sort(arr)
	minDiff := math.MaxInt
	for i, x := range arr[:len(arr)-1] {
		y := arr[i+1]
		diff := y - x
		if diff < minDiff {
			minDiff = diff
			ans = [][]int{{x, y}}
		} else if diff == minDiff {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
