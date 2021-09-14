package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func eliminateMaximum(dis, speed []int) int {
	n := len(dis)
	type pair struct{ d, spd int }
	a := make([]pair, n)
	for i, d := range dis {
		a[i] = pair{d, speed[i]}
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.d*b.spd < b.d*a.spd })
	for t, p := range a {
		if t*p.spd >= p.d {
			return t
		}
	}
	return n
}
