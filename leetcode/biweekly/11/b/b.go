package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func minAvailableDuration(a, b [][]int, duration int) (ans []int) {
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	sort.Slice(b, func(i, j int) bool { return b[i][0] < b[j][0] })
	i, j, n, m := 0, 0, len(a), len(b)
	for i < n && j < m {
		l, r := max(a[i][0], b[j][0]), min(a[i][1], b[j][1])
		if r-l >= duration {
			return []int{l, l + duration}
		}
		if a[i][1] < b[j][1] {
			i++
		} else {
			j++
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
