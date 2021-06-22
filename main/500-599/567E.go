package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type pair67 struct {
	v int
	d int64
}
type hp67 []pair67

func (h hp67) Len() int              { return len(h) }
func (h hp67) Less(i, j int) bool    { return h[i].d < h[j].d }
func (h hp67) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp67) Push(v interface{})   { *h = append(*h, v.(pair67)) }
func (h *hp67) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp67) push(v pair67)        { heap.Push(h, v) }
func (h *hp67) pop() pair67          { return heap.Pop(h).(pair67) }

func CF567E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, s, t, v, w, wt int
	Fscan(in, &n, &m, &s, &t)
	s--
	t--
	type edge struct {
		v, w, wt int
		isB      bool
	}
	es := make([]edge, m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	g2 := make([][]nb, n)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		es[i] = edge{v, w, wt, false}
		g[v] = append(g[v], nb{w, wt})
		g2[w] = append(g2[w], nb{v, wt})
	}

	dij := func(g [][]nb, st int) []int64 {
		dis := make([]int64, n)
		for i := range dis {
			dis[i] = 1e18
		}
		dis[st] = 0
		h := hp67{{st, 0}}
		for len(h) > 0 {
			vd := h.pop()
			v := vd.v
			if dis[v] < vd.d {
				continue
			}
			for _, e := range g[v] {
				w, wt := e.to, int64(e.wt)
				if newD := dis[v] + wt; newD < dis[w] {
					dis[w] = newD
					h.push(pair67{w, newD})
				}
			}
		}
		return dis
	}
	ds, dt := dij(g, s), dij(g2, t)

	// 将所有在最短路上的边组成一无向图 g3，跑割边，求出 YES。非割边就只能减 1 了，如果边权就是 1 则为 NO
	g3 := make([][]nb, n)
	for i, e := range es {
		if v, w := e.v, e.w; ds[v]+int64(e.wt)+dt[w] == ds[t] {
			g3[v] = append(g3[v], nb{w, i})
			g3[w] = append(g3[w], nb{v, i})
		}
	}
	dfn := make([]int, n)
	ts := 0
	var f func(int, int) int
	f = func(v, fid int) int {
		ts++
		dfn[v] = ts
		lowV := ts
		for _, e := range g3[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := f(w, e.wt)
				if lowW > dfn[v] {
					es[e.wt].isB = true
				}
				lowV = min(lowV, lowW)
			} else if e.wt != fid {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for v, t := range dfn {
		if t == 0 {
			f(v, -1)
		}
	}

	for _, e := range es {
		v, w, wt := e.v, e.w, int64(e.wt)
		if d := ds[v] + wt + dt[w] - ds[t]; d == 0 {
			if e.isB {
				Fprintln(out, "YES")
			} else if wt > 1 {
				Fprintln(out, "CAN 1")
			} else {
				Fprintln(out, "NO")
			}
		} else if wt > d+1 {
			Fprintln(out, "CAN", d+1)
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF567E(os.Stdin, os.Stdout) }
