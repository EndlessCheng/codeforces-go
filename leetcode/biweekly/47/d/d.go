package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func countPairs(n int, edges [][]int, queries []int) []int {
	deg := make([]int, n)
	for _, e := range edges {
		if e[0] > e[1] {
			e[0], e[1] = e[1], e[0]
		}
		e[0]--
		e[1]--
		deg[e[0]]++
		deg[e[1]]++
	}
	d := append([]int(nil), deg...)
	sort.Ints(d)
	ans := make([]int, len(queries))
	for q, low := range queries {
		for i, v := range d {
			ans[q] += i - sort.SearchInts(d[:i], low-v+1)
		}
		mp := map[[2]int]int{}
		for _, e := range edges {
			sum := deg[e[0]] + deg[e[1]]
			if sum <= low {
				continue
			}
			p := [2]int{e[0], e[1]}
			s, has := mp[p]
			if !has {
				mp[p] = sum
				s = sum
			}
			if s > low {
				mp[p]--
				if s-1 == low {
					ans[q]--
				}
			}
		}
	}
	return ans
}
