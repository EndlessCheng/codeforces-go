package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxConsecutive(bottom, top int, a []int) (ans int) {
	a = append(a, bottom-1, top+1)
	sort.Ints(a)
	for i := 1; i < len(a); i++ {
		ans = max(ans, a[i]-a[i-1]-1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
