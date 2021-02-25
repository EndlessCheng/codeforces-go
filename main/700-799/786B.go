package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair86 struct {
	v   int
	dis int64
}
type hp86 []pair86

func (h hp86) Len() int              { return len(h) }
func (h hp86) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp86) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp86) Push(v interface{})   { *h = append(*h, v.(pair86)) }
func (h *hp86) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp86) push(v pair86)        { heap.Push(h, v) }
func (h *hp86) pop() pair86          { return heap.Pop(h).(pair86) }

func CF786B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, st, tp, v, w, wt, l, r int
	Fscan(in, &n, &m, &st)
	d := 4 * n
	type nb struct{ to, wt int }
	g := make([][]nb, 2*d)
	leaf := make([]int, n+1)

	type seg []struct{ l, r int }
	t := make(seg, d)
	var build func(o, l, r int)
	build = func(o, l, r int) {
		t[o].l, t[o].r = l, r
		if l == r {
			leaf[l] = o
			g[o] = append(g[o], nb{o + d, 0})
			g[o+d] = append(g[o+d], nb{o, 0})
			return
		}
		lo, ro := o<<1, o<<1|1
		g[o] = append(g[o], nb{lo, 0}, nb{ro, 0})
		g[lo+d] = append(g[lo+d], nb{o + d, 0})
		g[ro+d] = append(g[ro+d], nb{o + d, 0})
		m := (l + r) >> 1
		build(lo, l, m)
		build(ro, m+1, r)
	}
	build(1, 1, n)

	var conn func(int)
	conn = func(o int) {
		if l <= t[o].l && t[o].r <= r {
			if tp == 2 {
				g[v] = append(g[v], nb{o, wt})
			} else {
				g[o+d] = append(g[o+d], nb{v, wt})
			}
			return
		}
		m := (t[o].l + t[o].r) >> 1
		if l <= m {
			conn(o << 1)
		}
		if m < r {
			conn(o<<1 | 1)
		}
	}
	for ; m > 0; m-- {
		if Fscan(in, &tp); tp == 1 {
			Fscan(in, &v, &w, &wt)
			g[leaf[v]] = append(g[leaf[v]], nb{leaf[w], wt})
		} else {
			Fscan(in, &v, &l, &r, &wt)
			v = leaf[v]
			conn(1)
		}
	}

	const inf int64 = 1e18
	dist := make([]int64, 2*d)
	for i := range dist {
		dist[i] = inf
	}
	st = leaf[st]
	dist[st] = 0
	q := hp86{{st, 0}}
	for len(q) > 0 {
		p := q.pop()
		v := p.v
		if dist[v] < p.dis {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			if newD := dist[v] + int64(e.wt); newD < dist[w] {
				dist[w] = newD
				q.push(pair86{w, newD})
			}
		}
	}
	for _, v := range leaf[1:] {
		if dist[v] < inf {
			Fprint(out, dist[v], " ")
		} else {
			Fprint(out, "-1 ")
		}
	}
}

//func main() { CF786B(os.Stdin, os.Stdout) }
