package main

import "slices"

// https://space.bilibili.com/206214
func sortTheStudents(score [][]int, k int) [][]int {
	slices.SortFunc(score, func(a, b []int) int { return b[k] - a[k] })
	return score
}
