package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func shortestDistanceColor(colors []int, queries [][]int) (ans []int) {
	pos := [4][]int{}
	for i, v := range colors {
		pos[v] = append(pos[v], i)
	}
	ans = make([]int, len(queries))
	for i, q := range queries {
		a := pos[q[1]]
		if a == nil {
			ans[i] = -1
			continue
		}
		cur := q[0]
		p := sort.SearchInts(a, cur)
		if p == 0 {
			ans[i] = a[0] - cur
		} else if p == len(a) {
			ans[i] = cur - a[len(a)-1]
		} else {
			ans[i] = min(a[p]-cur, cur-a[p-1])
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
