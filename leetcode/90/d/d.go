package main

import (
	"container/heap"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	type pair struct{ w, q int }
	ps := make([]pair, len(quality))
	for i, q := range quality {
		ps[i] = pair{wage[i], q}
	}
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.w*b.q < a.q*b.w })
	h, sum := &hp{make([]int, k)}, 0
	for i, p := range ps[:k] {
		h.IntSlice[i] = p.q
		sum += p.q
	}
	heap.Init(h)
	ans := float64(ps[k-1].w*sum) / float64(ps[k-1].q)
	for _, p := range ps[k:] {
		if q := p.q; q < h.IntSlice[0] {
			sum += q - h.IntSlice[0]
			ans = math.Min(ans, float64(p.w*sum)/float64(q))
			h.IntSlice[0] = q
			heap.Fix(h, 0)
		}
	}
	return ans
}
