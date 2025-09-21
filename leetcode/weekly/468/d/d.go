package main

import (
	"container/heap"
	"math/bits"
)

// https://space.bilibili.com/206214
type pair struct{ min, max int }

func op(a, b pair) pair {
	return pair{min(a.min, b.min), max(a.max, b.max)}
}

type ST [][]pair

func newST(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST, n)
	for i, x := range a {
		st[i] = make([]pair, sz)
		st[i][0] = pair{x, x}
	}
	for j := 1; j < sz; j++ {
		for i := range n - 1<<j + 1 {
			st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

// [l,r) 左闭右开
func (st ST) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	p := op(st[l][k], st[r-1<<k][k])
	return p.max - p.min
}

func maxTotalValue(nums []int, k int) (ans int64) {
	n := len(nums)
	st := newST(nums)
	h := make(hp, n)
	for i := range h {
		h[i] = tuple{st.query(i, n), i, n}
	}
	heap.Init(&h)

	for range k {
		t := heap.Pop(&h).(tuple)
		ans += int64(t.d)
		if t.l < t.r-1 {
			heap.Push(&h, tuple{st.query(t.l, t.r-1), t.l, t.r - 1})
		}
	}
	return
}

type tuple struct{ d, l, r int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d > h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
