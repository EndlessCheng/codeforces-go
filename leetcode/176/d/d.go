package main

import (
	"container/heap"
)

type hp []int64

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i] < h[j] } // > 为最大堆
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(int64)) }
func (h *hp) Pop() (v interface{}) { n := len(*h); *h, v = (*h)[:n-1], (*h)[n-1]; return }

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
	for h.Len() > 0 && (*h)[0] > 1 {
		top := heap.Pop(h).(int)
		d := 2*top - sum
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
