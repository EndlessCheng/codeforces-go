package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func largestSubmatrix(a [][]int) (ans int) {
	m := len(a[0])
	cnt := make([]int, m)
	for _, r := range a {
		for j, v := range r {
			if v == 0 {
				cnt[j] = 0
			} else {
				cnt[j]++
			}
		}
		c := append([]int(nil), cnt...)
		sort.Ints(c)
		for i, v := range c {
			ans = max(ans, (m-i)*v)
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
