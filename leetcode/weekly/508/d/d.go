package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
func minTimeMaxPower(n int, edges [][]int, power int, cost []int, source int, target int) []int64 {
	type edge struct{ to, t int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, t := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, t})
	}

	dis := make([]pair, n)
	for i := range dis {
		dis[i].d = math.MaxInt
	}
	dis[source] = pair{0, power}
	h := hp{{dis[source], source}}

	for len(h) > 0 {
		top := heap.Pop(&h).(tuple)
		d, rem, x := top.d, top.rem, top.x
		if x == target {
			return []int64{int64(d), int64(rem)}
		}
		if d > dis[x].d || d == dis[x].d && rem < dis[x].rem || rem < cost[x] {
			continue
		}
		rem -= cost[x]
		for _, e := range g[x] {
			y := e.to
			newD := pair{d + e.t, rem}
			if less(newD, dis[y]) {
				dis[y] = newD
				heap.Push(&h, tuple{newD, rem})
			}
		}
	}

	return []int64{-1, -1}
}

type pair struct{ d, rem int }

func less(a, b pair) bool {
	return a.d < b.d || a.d == b.d && a.rem > b.rem
}

// 最短路长度, 剩余电量, 节点编号
type tuple struct {
	pair
	x int
}
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i].pair, h[j].pair) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
