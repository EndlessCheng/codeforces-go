package main

import "sort"

func minTaps(n int, ranges []int) (ans int) {
	type pair struct{ l, r int }
	ps := make([]pair, 0, n+1)
	for i, r := range ranges {
		if r > 0 {
			ps = append(ps, pair{i - r, i + r})
		}
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].l < ps[j].l })
	vis := make([]bool, n+1)
	j := 0
	for i, v := range vis {
		if v {
			continue
		}
		r := -1
		for ; j < len(ps); j++ {
			if ps[j].l > i {
				break
			}
			if ps[j].r > r {
				r = ps[j].r
			}
		}
		if r == -1 {
			return -1
		}
		if r >= n {
			r = n + 1
		}
		for k := i; k < r; k++ {
			vis[k] = true
		}
		ans++
	}
	return
}
