package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

func maxPerformance(n int, speed []int, efficiency []int, k int) (ans int) {
	type pair struct{ eff, spd int }
	ps := make([]pair, n)
	for i := range ps {
		ps[i] = pair{efficiency[i], speed[i]}
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.eff > b.eff || a.eff == b.eff && a.spd > b.spd })
	sumSpd := 0
	h := &hp{}
	for _, p := range ps {
		if len(h.IntSlice) >= k {
			v := heap.Pop(h).(int)
			sumSpd -= v
		}
		heap.Push(h, p.spd)
		sumSpd += p.spd
		if val := p.eff * sumSpd; val > ans {
			ans = val
		}
	}
	return ans % (1e9 + 7)
}
