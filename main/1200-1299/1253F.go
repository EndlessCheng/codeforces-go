package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1253F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, k, q int
	Fscan(in, &n, &m, &k, &q)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
		es[i] = edge{v, w, wt}
	}

	dis := make([]int, n)
	for i := k; i < n; i++ {
		dis[i] = 1e18
	}
	h := make(hp53, k)
	for i := range h {
		h[i].v = i
	}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair53)
		if p.d > dis[p.v] {
			continue
		}
		for _, e := range g[p.v] {
			w := e.to
			newD := p.d + e.wt
			if newD < dis[w] {
				dis[w] = newD
				heap.Push(&h, pair53{newD, w})
			}
		}
	}

	for i, e := range es {
		es[i].wt += dis[e.v] + dis[e.w]
	}
	slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

	type query struct{ to, i int }
	qs := make([][]query, n)
	for i := range q {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		qs[v] = append(qs[v], query{w, i})
		qs[w] = append(qs[w], query{v, i})
	}

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, q)
	for _, e := range es {
		v, w := find(e.v), find(e.w)
		if v == w {
			continue
		}
		if len(qs[v]) > len(qs[w]) {
			v, w = w, v
		}
		for _, q := range qs[v] {
			if ans[q.i] > 0 {
				continue
			}
			if find(q.to) == w {
				ans[q.i] = e.wt
			} else {
				qs[w] = append(qs[w], q)
			}
		}
		fa[v] = w
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1253F(bufio.NewReader(os.Stdin), os.Stdout) }

type pair53 struct{ d, v int }
type hp53 []pair53
func (h hp53) Len() int           { return len(h) }
func (h hp53) Less(i, j int) bool { return h[i].d < h[j].d }
func (h hp53) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp53) Push(v any)        { *h = append(*h, v.(pair53)) }
func (h *hp53) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
