package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF733F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, money int
	Fscan(in, &n, &m)
	es := make([]struct{ v, w, wt, c, i int }, m)
	for i := range es {
		Fscan(in, &es[i].wt)
		es[i].i = i
	}
	for i := range es {
		Fscan(in, &es[i].c)
	}
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w)
		es[i].v--
		es[i].w--
	}
	Fscan(in, &money)
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	s := int64(0)
	type nb struct{ to, wt, eid int }
	g := make([][]nb, n)
	for i, e := range es {
		v, w, wt, eid := e.v, e.w, e.wt, e.i
		if fv, fw := find(v), find(w); fv != fw {
			s += int64(wt)
			fa[fv] = fw
			g[v] = append(g[v], nb{w, wt, eid})
			g[w] = append(g[w], nb{v, wt, eid})
			es[i].wt = -es[i].wt
		}
	}

	const mx = 18
	type pair struct{ p, max, eid int }
	pa := make([][mx]pair, n)
	dep := make([]int, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		pa[v][0].p = p
		dep[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].max = e.wt
				pa[w][0].eid = e.eid
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1].p = pp.p
				if p.max > pp.max {
					pa[v][i+1].max = p.max
					pa[v][i+1].eid = p.eid
				} else {
					pa[v][i+1].max = pp.max
					pa[v][i+1].eid = pp.eid
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}
	maxWt := func(v, w int) (mxWt, eid int) {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (dep[w]-dep[v])>>i&1 > 0 {
				p := pa[w][i]
				w = p.p
				if p.max > mxWt {
					mxWt, eid = p.max, p.eid
				}
			}
		}
		if v == w {
			return
		}
		for i := mx - 1; i >= 0; i-- {
			if p, q := pa[v][i], pa[w][i]; p.p != q.p {
				v, w = p.p, q.p
				if p.max > mxWt {
					mxWt, eid = p.max, p.eid
				}
				if q.max > mxWt {
					mxWt, eid = q.max, q.eid
				}
			}
		}
		if p := pa[v][0]; p.max > mxWt {
			mxWt, eid = p.max, p.eid
		}
		if p := pa[w][0]; p.max > mxWt {
			mxWt, eid = p.max, p.eid
		}
		return
	}

	mxDec, ori, cur := -1, -1, 0
	for _, e := range es {
		dec := money / e.c
		if e.wt > 0 {
			mxWt, eid := maxWt(e.v, e.w)
			dec = mxWt - (e.wt - dec)
			if dec > mxDec {
				mxDec, ori, cur = dec, eid, e.i
			}
		} else {
			if dec > mxDec {
				mxDec, ori, cur = dec, -1, e.i
			}
		}
	}
	Fprintln(out, s-int64(mxDec))
	for _, e := range es {
		if e.i == ori || e.i != cur && e.wt > 0 {
			continue
		}
		wt := e.wt
		if wt < 0 {
			wt = -wt
		}
		if e.i == cur {
			wt -= money / e.c
		}
		Fprintln(out, e.i+1, wt)
	}
}

//func main() { CF733F(os.Stdin, os.Stdout) }
