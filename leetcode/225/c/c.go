package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func kthLargestValue(a [][]int, k int) int {
	n, m := len(a), len(a[0])
	b := make([]int, 0, n*m)
	s := make([][]int, n+1)
	s[0] = make([]int, m+1)
	for i, row := range a {
		s[i+1] = make([]int, m+1)
		for j, v := range row {
			s[i+1][j+1] = s[i+1][j] ^ s[i][j+1] ^ s[i][j] ^ v
			b = append(b, s[i+1][j+1])
		}
	}
	sort.Ints(b)
	return b[n*m-k]
}
