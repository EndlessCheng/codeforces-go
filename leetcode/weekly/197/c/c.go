package main

import (
	"container/heap"
	"runtime/debug"
)

// github.com/EndlessCheng/codeforces-go
func init() { debug.SetGCPercent(-1) }

type pair struct {
	dis float64
	v   int
}
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis > h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }

func maxProbability(n int, edges [][]int, succProb []float64, st, end int) (ans float64) {
	type neighbor struct {
		to     int
		weight float64
	}
	g := make([][]neighbor, n)
	for i, e := range edges {
		v, w, weight := e[0], e[1], succProb[i]
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	dist := make([]float64, n)
	dist[st] = 1
	h := hp{{1, st}}
	for len(h) > 0 {
		p := h.pop()
		d, v := p.dis, p.v
		if dist[v] < d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := d * e.weight; newD > dist[w] {
				dist[w] = newD
				h.push(pair{newD, w})
			}
		}
	}
	return dist[end]
}
