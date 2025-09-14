package main

import "math"

// https://space.bilibili.com/206214
func earliestTime(tasks [][]int) int {
	ans := math.MaxInt
	for _, t := range tasks {
		ans = min(ans, t[0]+t[1])
	}
	return ans
}
