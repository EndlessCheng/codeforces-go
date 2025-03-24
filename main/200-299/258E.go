package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg58 []struct{ l, r, s, minCover, todo int }

func (t seg58) maintain(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	mn := min(lo.minCover, ro.minCover)
	t[o].minCover = mn
	t[o].s = 0
	if lo.minCover == mn {
		t[o].s = lo.s
	}
	if ro.minCover == mn {
		t[o].s += ro.s
	}
}

func (t seg58) do(o, v int) {
	t[o].minCover += v
	t[o].todo += v
}

func (t seg58) spread(o int) {
	v := t[o].todo
	if v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg58) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].s = 1
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg58) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func cf258E(in io.Reader, out io.Writer) {
	var n, m, dfn int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	a := make([]struct{ l, r int }, n)
	dfnToV := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		a[v].l = dfn
		dfnToV[dfn] = v
		dfn++
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
			}
		}
		a[v].r = dfn - 1
	}
	dfs(0, -1)

	type event struct{ lx, rx, delta int }
	events := make([][]event, n+1)
	add := func(lx, ly, rx, ry int) {
		events[ly] = append(events[ly], event{lx, rx, 1})
		events[ry+1] = append(events[ry+1], event{lx, rx, -1})
	}
	for range m {
		var v, w int
		Fscan(in, &v, &w)
		p := a[v-1]
		q := a[w-1]
		add(p.l, p.l, p.r, p.r)
		add(p.l, q.l, p.r, q.r)
		add(q.l, p.l, q.r, p.r)
		add(q.l, q.l, q.r, q.r)
	}

	ans := make([]any, n)
	t := make(seg58, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	for i, es := range events[:n] {
		for _, e := range es {
			t.update(1, e.lx, e.rx, e.delta)
		}
		res := n
		if t[1].minCover == 0 {
			res -= t[1].s
		}
		ans[dfnToV[i]] = max(res-1, 0)
	}
	Fprintln(out, ans...)
}

//func main() { cf258E(bufio.NewReader(os.Stdin), os.Stdout) }
