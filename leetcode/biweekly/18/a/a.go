package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func arrayRankTransform(a []int) []int {
	if len(a) == 0 {
		return a
	}
	type pair struct{ v, i int }
	ps := make([]pair, len(a))
	for i, v := range a {
		ps[i] = pair{v, i}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v })
	k := 1
	a[ps[0].i] = k
	for i := 1; i < len(ps); i++ {
		if ps[i].v != ps[i-1].v {
			k++
		}
		a[ps[i].i] = k
	}
	return a
}
