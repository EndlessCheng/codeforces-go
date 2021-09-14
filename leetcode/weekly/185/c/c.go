package main

import "sort"

func minNumberOfFrogs(s string) (ans int) {
	ans = -1
	n := len(s)
	if n%5 != 0 {
		return
	}

	pos := [5][]int{}
	for i, b := range s {
		for j, b2 := range "croak" {
			if b == b2 {
				pos[j] = append(pos[j], i)
				break
			}
		}
	}
	for _, p := range pos {
		if len(p) != n/5 {
			return
		}
	}

	stops := []int{}
	for j := range pos[0] {
		ps := make([]int, 5)
		for i, p := range pos {
			ps[i] = p[j]
		}
		if !sort.IntsAreSorted(ps) {
			return
		}
		if len(stops) > 0 && stops[0] <= ps[0] {
			stops = stops[1:]
		}
		stops = append(stops, ps[4]+1)
	}
	return len(stops)
}
