package main

import (
	"container/heap"
	"math"
)

// https://space.bilibili.com/206214
type Graph [][]pair

func Constructor(n int, edges [][]int) Graph {
	g := make(Graph, n) // 邻接表
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], pair{e[1], e[2]})
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]] = append(g[e[0]], pair{e[1], e[2]})
}

func (g Graph) ShortestPath(start, end int) int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		x, d := p.x, p.d
		if x == end {
			return d
		}
		if d > dis[x] {
			continue
		}
		for _, e := range g[x] {
			y, w := e.x, e.d
			newD := d + w
			if newD < dis[y] {
				dis[y] = newD
				heap.Push(&h, pair{y, newD})
			}
		}
	}
	return -1
}

type pair struct{ x, d int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
