package main

import "sort"

// https://space.bilibili.com/206214
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	ans := make([]int, len(obstacles))
	g := []int{}
	for i, x := range obstacles {
		p := sort.SearchInts(g, x+1)
		if p < len(g) {
			g[p] = x
		} else {
			g = append(g, x)
		}
		ans[i] = p + 1
	}
	return ans
}
