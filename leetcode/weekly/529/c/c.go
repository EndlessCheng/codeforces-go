package main

import "container/heap"

// https://space.bilibili.com/206214
func earliestActivationTimes(n int, edges [][]int, requiredSignals []int) []int64 {
	type edge struct{ to, delay int }
	g := make([][]edge, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], edge{e[1], e[2]})
	}

	ans := make([]int64, n)
	for i := range ans {
		ans[i] = -1
	}
	h := &hp{{}}
	requiredSignals[0] = 1

	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		t := top.t
		x := top.x
		requiredSignals[x]--
		if requiredSignals[x] != 0 { // x 尚未激活或者已经激活
			continue
		}
		ans[x] = t
		for _, e := range g[x] {
			heap.Push(h, pair{t + int64(e.delay), e.to})
		}
	}

	return ans
}

type pair struct {
	t int64
	x int
}
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].t < h[j].t }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
