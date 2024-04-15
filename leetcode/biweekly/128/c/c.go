package main

import (
	"container/heap"
)

// https://space.bilibili.com/206214
func minimumTime(n int, edges [][]int, disappear []int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	h := hp{{}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		dx := p.dis
		x := p.x
		if dx > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y := e.to
			newD := dx + e.wt
			if newD < disappear[y] && (dis[y] < 0 || newD < dis[y]) {
				dis[y] = newD
				heap.Push(&h, pair{newD, y})
			}
		}
	}
	return dis
}

type pair struct{ dis, x int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

// 居然有卡 SPFA 的数据
func minimumTimeSPFA_TLE(n int, edges [][]int, disappear []int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		x, y, wt := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, wt})
		g[y] = append(g[y], edge{x, wt})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	dis[0] = 0
	q := []int{0}
	inQ := make([]bool, n)
	inQ[0] = true
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		inQ[v] = false
		for _, e := range g[v] {
			w := e.to
			newD := dis[v] + e.wt
			if newD < disappear[w] && (dis[w] < 0 || newD < dis[w]) {
				dis[w] = newD
				if !inQ[w] {
					inQ[w] = true
					q = append(q, w)
				}
			}
		}
	}
	return dis
}
