package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maxSumRangeQuery(a []int, requests [][]int) (ans int) {
	n := len(a)
	d := make([]int, len(a)+1)
	for _, r := range requests {
		d[r[0]]++
		d[r[1]+1]--
	}
	c := make([]int, n)
	s := 0
	for i, v := range d[:n] {
		s += v
		c[i] = s
	}
	sort.Ints(c)
	sort.Ints(a)
	for i, v := range a {
		ans += c[i] * v
	}
	return ans % (1e9 + 7)
}
