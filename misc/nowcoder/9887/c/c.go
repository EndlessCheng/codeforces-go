package main

import (
	"container/heap"
	"math"
)

// github.com/EndlessCheng/codeforces-go
const mod, mx int = 1e9 + 7, 1e3

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

type vd struct {
	v, m int
	d    float64
}
type hp []vd

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(vd)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp) push(v vd)            { heap.Push(h, v) }
func (h *hp) pop() vd              { return heap.Pop(h).(vd) }

func minDist(n, _, s, t int, es [][]int) int {
	L := [mx + 1]float64{}
	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		L[i] = L[i-1] + math.Log(float64(i))
		F[i] = F[i-1] * i % mod
	}
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}

	g := make([][]vd, n+1)
	for _, e := range es {
		v, w, a, b := e[0], e[1], e[2], e[3]
		d := L[a] - L[b] - L[a-b]
		m := F[a] * invF[b] % mod * invF[a-b] % mod
		g[v] = append(g[v], vd{w, m, d})
		g[w] = append(g[w], vd{v, m, d})
	}
	dis := make([]vd, n+1)
	for i := range dis {
		dis[i] = vd{0, 1, 1e9}
	}
	dis[s].d = 0
	q := hp{{s, 1, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dis[v].d < p.d {
			continue
		}
		for _, e := range g[v] {
			w := e.v
			if d := dis[v].d + e.d; d < dis[w].d {
				dis[w] = vd{w, (dis[v].m * e.m) % mod, d}
				q.push(dis[w])
			}
		}
	}
	return dis[t].m
}
