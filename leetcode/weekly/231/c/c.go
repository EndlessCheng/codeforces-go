package main

import "container/heap"

// github.com/EndlessCheng/codeforces-go
type vdPair struct{ v, d int }
type hp []vdPair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v vdPair)        { heap.Push(h, v) }
func (h *hp) pop() vdPair          { return heap.Pop(h).(vdPair) }

type nb struct{ to, wt int }

func dijkstra(g [][]nb, st int) []int {
	const inf int = 1e18
	dist := make([]int, len(g))
	for i := range dist {
		dist[i] = inf
	}
	dist[st] = 0
	q := hp{{st, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dist[v] < p.d {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + e.wt; newD < dist[w] {
				dist[w] = newD
				q.push(vdPair{w, newD})
			}
		}
	}
	return dist
}

func countRestrictedPaths(n int, edges [][]int) (ans int) {
	const mod int = 1e9 + 7
	g := make([][]nb, n)
	for _, e := range edges {
		v, w, wt := e[0]-1, e[1]-1, e[2]
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}
	dist := dijkstra(g, n-1)

	// 递推写法是排序 dist，按其升序转移
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	var f func(int) int
	f = func(v int) (res int) {
		if v == n-1 {
			return 1
		}
		dv := &dp[v]
		if *dv != -1 {
			return *dv
		}
		defer func() { *dv = res }()
		for _, e := range g[v] {
			if w := e.to; dist[w] < dist[v] {
				res += f(w)
			}
		}
		return res % mod
	}
	return f(0) % mod
}
