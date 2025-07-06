package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func minTime(n int, edges [][]int) int {
	type edge struct{ to, start, end int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], edge{y, e[2], e[3]})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		d := p.d
		x := p.x
		if d > dis[x] {
			continue
		}
		if x == n-1 {
			return d
		}
		for _, e := range g[x] {
			y := e.to
			newD := max(d, e.start) + 1
			if newD-1 <= e.end && newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}
	return -1
}

type pair struct{ d, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
