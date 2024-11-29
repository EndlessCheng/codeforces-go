package main

import (
	"container/heap"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func maxSpending(values [][]int) (ans int64) {
	m, n := len(values), len(values[0])
	a := make([]int, 0, m*n) // 预分配空间
	for _, row := range values {
		a = append(a, row...)
	}
	slices.Sort(a)

	for i, x := range a {
		ans += int64(x) * int64(i+1)
	}
	return
}

func maxSpending2(values [][]int) (ans int64) {
	m, n := len(values), len(values[0])
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	h := &hp{idx, values}
	heap.Init(h)

	for d := 1; d <= m*n; d++ {
		a := values[idx[0]]
		ans += int64(a[len(a)-1]) * int64(d)
		if len(a) > 1 {
			values[idx[0]] = a[:len(a)-1]
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return
}

type hp struct {
	sort.IntSlice
	values [][]int
}

func (h hp) Less(i, j int) bool {
	a, b := h.values[h.IntSlice[i]], h.values[h.IntSlice[j]]
	return a[len(a)-1] < b[len(b)-1]
}
func (hp) Push(any)        {}
func (h *hp) Pop() (_ any) { a := h.IntSlice; h.IntSlice = a[:len(a)-1]; return }
