package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type vd51 struct {
	v   int
	dis int64
}
type hp51 []vd51

func (h hp51) Len() int              { return len(h) }
func (h hp51) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp51) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp51) Push(v interface{})   { *h = append(*h, v.(vd51)) }
func (h *hp51) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp51) push(v vd51)          { heap.Push(h, v) }
func (h *hp51) pop() vd51            { return heap.Pop(h).(vd51) }

func CF1051F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n, m, q, v, w, wt int
	Fscan(in, &n, &m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	spV := map[int]bool{}
	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	dep2 := make([]int64, n)
	var f func(int, int)
	f = func(v, fa int) {
		pa[v][0] = fa
		for _, e := range g[v] {
			w := e.to
			if dep[w] == 0 {
				dep[w] = dep[v] + 1
				dep2[w] = dep2[v] + int64(e.wt)
				f(w, v)
			} else if w != fa {
				spV[v] = true
				spV[w] = true
			}
		}
	}
	f(0, -1)

	spDis := make([][]int64, 0, len(spV))
	for v := range spV {
		dis := make([]int64, n)
		for i := range dis {
			dis[i] = 1e18
		}
		dis[v] = 0
		h := hp51{{v, 0}}
		for len(h) > 0 {
			vd := h.pop()
			v := vd.v
			if dis[v] < vd.dis {
				continue
			}
			for _, e := range g[v] {
				w := e.to
				if newD := dis[v] + int64(e.wt); newD < dis[w] {
					dis[w] = newD
					h.push(vd51{w, newD})
				}
			}
		}
		spDis = append(spDis, dis)
	}

	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for i := 0; i < mx; i++ {
			if (dep[v]-d)>>i&1 > 0 {
				v = pa[v][i]
			}
		}
		return v
	}
	_lca := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &v, &w)
		v--
		w--
		d := dep2[v] + dep2[w] - dep2[_lca(v, w)]*2
		for _, dis := range spDis {
			d = min(d, dis[v]+dis[w])
		}
		Fprintln(out, d)
	}
}

//func main() { CF1051F(os.Stdin, os.Stdout) }
