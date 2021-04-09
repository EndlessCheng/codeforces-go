package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func countPairs(n int, es [][]int, qs []int) []int {
	deg := make([]int, n)
	type edge struct{ v, w int }
	cntE := map[edge]int{}
	for _, e := range es {
		v, w := e[0]-1, e[1]-1
		if v > w {
			v, w = w, v
		}
		deg[v]++
		deg[w]++
		cntE[edge{v, w}]++
	}

	d := append([]int(nil), deg...)
	sort.Ints(d)
	ans := make([]int, len(qs))
	for q, low := range qs {
		low++
		for e, c := range cntE {
			if s := deg[e.v] + deg[e.w]; low <= s && s-c < low {
				ans[q]--
			}
		}
		for i, c := range d {
			ans[q] += i - sort.SearchInts(d[:i], low-c)
		}
	}
	return ans
}
