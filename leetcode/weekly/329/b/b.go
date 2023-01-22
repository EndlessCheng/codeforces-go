package main

import "sort"

// https://space.bilibili.com/206214
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool { return score[i][k] > score[j][k] })
	return score
}
