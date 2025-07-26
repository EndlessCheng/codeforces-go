package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf827D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt, i int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
		es[i].v--
		es[i].w--
		es[i].i = i
	}
	slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

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

	type nb struct{ to, wt, i int }
	g := make([][]nb, n)
	for i, e := range es {
		v, w := e.v, e.w
		fv, fw := find(v), find(w)
		if fv != fw {
			fa[fv] = fw
			g[v] = append(g[v], nb{w, e.wt, e.i})
			g[w] = append(g[w], nb{v, e.wt, e.i})
			es[i].wt = -1
		}
	}

	const mx = 18
	type pair struct{ p, maxWt int }
	pa := make([][mx]pair, n)
	paI := make([]int, n)
	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, p int) {
		pa[v][0].p = p
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].maxWt = e.wt
				paI[w] = e.i
				dep[w] = dep[v] + 1
				dfs(w, v)
			}
		}
	}
	dfs(0, -1)
	for i := range mx - 1 {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1] = pair{pp.p, max(p.maxWt, pp.maxWt)}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}
	getLCA := func(v, w int) (lca, maxWt int) {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for k := dep[w] - dep[v]; k > 0; k &= k - 1 {
			p := pa[w][bits.TrailingZeros(uint(k))]
			maxWt = max(maxWt, p.maxWt)
			w = p.p
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				pv, pw := pa[v][i], pa[w][i]
				if pv.p != pw.p {
					maxWt = max(maxWt, pv.maxWt, pw.maxWt)
					v, w = pv.p, pw.p
				}
			}
			maxWt = max(maxWt, pa[v][0].maxWt, pa[w][0].maxWt)
			v = pa[v][0].p
		}
		lca = v
		return
	}

	ans := make([]int, m)
	for i := range ans {
		ans[i] = -1
	}
	for i := range fa {
		fa[i] = i
	}
	f := func(v, tar, wt int) {
		for v = find(v); v != tar; v = find(pa[v][0].p) {
			ans[paI[v]] = wt
			fa[v] = tar
		}
	}
	for _, e := range es {
		if e.wt < 0 {
			continue
		}
		v, w := e.v, e.w
		lca, maxWt := getLCA(v, w)
		ans[e.i] = maxWt - 1
		tar := find(lca)
		f(v, tar, e.wt-1)
		f(w, tar, e.wt-1)
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf827D(bufio.NewReader(os.Stdin), os.Stdout) }
