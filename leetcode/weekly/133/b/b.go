package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func allCellsDistOrder(r, c, r0, c0 int) (ans [][]int) {
	d := func(p []int) int { return abs(p[0]-r0) + abs(p[1]-c0) }
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans = append(ans, []int{i, j})
		}
	}
	sort.Slice(ans, func(i, j int) bool { return d(ans[i]) < d(ans[j]) })
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
