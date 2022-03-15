package main

import (
	"container/heap"
	"math"
)

// github.com/EndlessCheng/codeforces-go
type edge struct{ to, wt int }
func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = math.MaxInt64 / 3
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if p.dis > dis[v] {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dis[v] + e.wt; newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, pair{w, newD})
			}
		}
	}
	return dis
}

func minimumWeight(n int, edges [][]int, src1, src2, dest int) int64 {
	g := make([][]edge, n)
	rg := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]
		g[v] = append(g[v], edge{w, wt})
		rg[w] = append(rg[w], edge{v, wt})
	}

	d1 := dijkstra(g, src1)
	d2 := dijkstra(g, src2)
	d3 := dijkstra(rg, dest)

	ans := math.MaxInt64 / 3
	for x := 0; x < n; x++ {
		ans = min(ans, d1[x]+d2[x]+d3[x])
	}
	if ans < math.MaxInt64 / 3 {
		return int64(ans)
	}
	return -1
}

type pair struct{ v, dis int }
type hp []pair
func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func min(a, b int) int { if a > b { return b }; return a }
