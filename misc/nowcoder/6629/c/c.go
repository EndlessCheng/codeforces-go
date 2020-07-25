package main

import "container/heap"

type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis } // > 权值最大
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v pair)          { heap.Push(h, v) }
func (h *hp) pop() pair            { return heap.Pop(h).(pair) }

// github.com/EndlessCheng/codeforces-go
func Length(vs []int, ws []int, path [][]int, n int) int {
	const inf int = 1e18
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(vs) > len(ws) {
		vs, ws = ws, vs
	}
	type neighbor struct{ to, weight int }
	g := make([][]neighbor, n)
	for _, e := range path {
		v, w, weight := e[0]-1, e[1]-1, e[2]
		g[v] = append(g[v], neighbor{w, weight})
		g[w] = append(g[w], neighbor{v, weight})
	}

	dij := func(st int) int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		h := hp{{st, 0}}
		for len(h) > 0 {
			p := h.pop()
			v, d := p.v, p.dis
			if dist[v] < d {
				continue
			}
			for _, e := range g[v] {
				w := e.to
				if newD := d + e.weight; newD < dist[w] {
					dist[w] = newD
					h.push(pair{w, newD})
				}
			}
		}
		minD := inf
		for _, w := range ws {
			minD = min(minD, dist[w-1])
		}
		return minD
	}
	ans := inf
	for _, v := range vs {
		ans = min(ans, dij(v-1))
	}
	if ans == inf {
		ans = -1
	}
	return ans
}
