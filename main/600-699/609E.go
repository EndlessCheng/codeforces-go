package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF609E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, m, v, w, wt int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt, i int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		es[i] = edge{v - 1, w - 1, wt, i}
	}
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
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i, e := range es {
		v, w, wt := e.v, e.w, e.wt
		if fv, fw := find(v), find(w); fv != fw {
			s += int64(wt)
			fa[fv] = fw
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
			es[i].wt = 0
		}
	}

	const mx = 18
	type pair struct{ p, max int }
	pa := make([][mx]pair, n)
	dep := make([]int, n)
	var f func(v, p, d int)
	f = func(v, p, d int) {
		pa[v][0].p = p
		dep[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].max = e.wt
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 0)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1] = pair{pp.p, max(p.max, pp.max)}
			} else {
				pa[v][i+1] = pair{-1, 0}
			}
		}
	}
	maxWt := func(v, w int) (res int) {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (dep[w]-dep[v])>>i&1 > 0 {
				p := pa[w][i]
				w = p.p
				res = max(res, p.max)
			}
		}
		if v == w {
			return
		}
		for i := mx - 1; i >= 0; i-- {
			if p, q := pa[v][i], pa[w][i]; p.p != q.p {
				v, w = p.p, q.p
				res = max(res, max(p.max, q.max))
			}
		}
		return max(res, max(pa[v][0].max, pa[w][0].max))
	}

	ans := make([]int64, m)
	for _, e := range es {
		if e.wt > 0 {
			ans[e.i] = int64(e.wt - maxWt(e.v, e.w))
		}
	}
	for _, v := range ans {
		Fprintln(out, s+v)
	}
}

//func main() { CF609E(os.Stdin, os.Stdout) }
