package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func diagonalSort(a [][]int) [][]int {
	n, m := len(a), len(a[0])
	for s := 1; s < n+m; s++ {
		l := max(0, m-s)
		r := min(m-1, m-s+n-1)
		b := []int{}
		for j := l; j <= r; j++ {
			i := s + j - m
			b = append(b, a[i][j])
		}
		sort.Ints(b)
		for j := l; j <= r; j++ {
			i := s + j - m
			a[i][j] = b[j-l]
		}
	}
	return a
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
