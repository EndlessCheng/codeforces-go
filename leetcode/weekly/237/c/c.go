package main

import (
	"container/heap"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type pair struct{ l, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.l < b.l || a.l == b.l && a.i < b.i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }

func getOrder(a [][]int) (ans []int) {
	for i := range a {
		a[i] = append(a[i], i)
	}
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })
	h := &hp{}
	for i, cur, n := 0, 0, len(a); i < n; {
		if h.Len() > 0 {
			p := h.pop()
			ans = append(ans, p.i)
			cur += p.l
		}
		if h.Len() == 0 && cur < a[i][0] {
			cur = a[i][0]
		}
		for ; i < n && a[i][0] <= cur; i++ {
			h.push(pair{a[i][1], a[i][2]})
		}
	}
	for h.Len() > 0 {
		ans = append(ans, h.pop().i)
	}
	return
}
