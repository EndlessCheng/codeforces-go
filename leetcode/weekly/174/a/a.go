package main

import "sort"

func kWeakestRows(mat [][]int, k int) (ans []int) {
	type pair struct{ x, y int }
	ps := make([]pair, len(mat))
	for i, mi := range mat {
		for _, v := range mi {
			ps[i].x += v
		}
		ps[i].y = i
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.x < b.x || a.x == b.x && a.y < b.y })
	for _, p := range ps[:k] {
		ans = append(ans, p.y)
	}
	return
}
