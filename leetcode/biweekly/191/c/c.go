package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
const mx = 100_001
var dis [mx]int

func init() {
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = -1
	h := &hp{{-1, 0}}
	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		d, s := top.dis, top.v
		if d > dis[s] {
			continue
		}
		t, i := 1, 1
		for s+t < mx {
			w := s + t
			newD := d + i + 1
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(h, pair{newD, w})
			}
			i++
			t += i
		}
	}
}

func minDays(n int) int {
	return dis[n]
}

type pair struct{ dis, v int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
