package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

func isPossible(a []int) (ans bool) {
	if len(a) == 1 {
		return a[0] == 1
	}
	sum := 0
	h := &hp{}
	for _, v := range a {
		sum += v
		heap.Push(h, v)
	}
	for h.Len() > 0 && h.IntSlice[0] > 1 {
		top := heap.Pop(h).(int)
		d := 2*top - sum // TODO: 虽然过了但是取模更好，这样不会在 [1,1e9] 这样的数据下 TLE
		if d <= 0 {
			return
		}
		if d > 1 {
			heap.Push(h, d)
		}
		sum = top
	}
	return true
}

// 另一种思路
func isPossible2(a []int) bool {
	n := len(a)
	if n == 1 {
		return a[0] == 1
	}
	sort.Ints(a)
	sum := 0
	prev := 1
	for i, v := range a {
		if v > 1 && (v == prev || v < sum+n-i) || (v-prev)%(n-1) != 0 {
			return false
		}
		sum += v
		prev = v
	}
	return true
}
