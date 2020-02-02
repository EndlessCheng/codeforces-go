package main

import "sort"

func kWeakestRows(mat [][]int, k int) (ans []int) {
	type pair struct{ cnt1, i int }
	ps:= make([]pair, len(mat))
	for i, mi := range mat {
		for _, v := range mi {
			if v == 1 {
				ps[i].cnt1++
			}
		}
		ps[i].i = i
	}
	sort.Slice(ps, func(i, j int) bool {
		pi, pj := ps[i], ps[j]
		return pi.cnt1 < pj.cnt1 || pi.cnt1 == pj.cnt1 && pi.i < pj.i
	})
	for _, p := range ps[:k] {
		ans = append(ans, p.i)
	}
	return
}
