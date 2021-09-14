package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxAscendingSum(a []int) (ans int) {
	for l := range a {
		for r := l + 1; r <= len(a); r++ {
			if b := a[l:r]; sort.SliceIsSorted(b, func(i, j int) bool { return b[i] <= b[j] }) {
				s := 0
				for _, v := range b {
					s += v
				}
				ans = max(ans, s)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
