package main

import "container/heap"

// https://space.bilibili.com/206214
func supplyWagon(a []int) []int {
	m := len(a) / 2
	n := len(a)
	left := make([]int, n+2)
	right := make([]int, n+1)
	h := make(hp, n)
	for i := 1; i <= n; i++ {
		left[i], right[i] = i-1, i+1
		h[i-1] = pair{a[i-1]+a[i],i}
	}
	heap.Init(&h)
	right[0] = 1
	left[n+1] = n
	del := func(i int) {
		l, r := left[i], right[i]
		right[l] = r
		left[r] = l
	}

	for len(h) > m {

	}

	return a
}

type pair struct{ s, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.s < b.s || a.s == b.s && a.i < b.i }
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair) { heap.Push(h, v) }
func (h *hp) pop() pair   { return heap.Pop(h).(pair) }
func (h hp) empty() bool   { return len(h) == 0 }
func (h hp) top() pair    { return h[0] }
func (h *hp) init()        { heap.Init(h) }

// 需保证 h 非空
func (h *hp) replace(v pair) pair { top := (*h)[0]; (*h)[0] = v; heap.Fix(h, 0); return top }
