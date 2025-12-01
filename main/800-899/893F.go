package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type node93 struct {
	lo, ro *node93
	mn     int
}

func build93(l, r int) *node93 {
	o := &node93{mn: 1e9}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build93(l, m)
	o.ro = build93(m+1, r)
	return o
}

func (o node93) update(l, r, i, val int) *node93 {
	if l == r {
		o.mn = val
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		o.lo = o.lo.update(l, m, i, val)
	} else {
		o.ro = o.ro.update(m+1, r, i, val)
	}
	o.mn = min(o.lo.mn, o.ro.mn)
	return &o
}

func (o *node93) query(l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return o.mn
	}
	m := (l + r) >> 1
	if qr <= m {
		return o.lo.query(l, m, ql, qr)
	}
	if m < ql {
		return o.ro.query(m+1, r, ql, qr)
	}
	return min(o.lo.query(l, m, ql, qr), o.ro.query(m+1, r, ql, qr))
}

func cf893F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, rt, dfn, m, x, k, ans int
	Fscan(in, &n, &rt)
	rt--
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	nodes := make([]struct{ l, r int }, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		nodes[v].l = dfn
		dfn++
		for _, w := range g[v] {
			if w != fa {
				dfs(w, v)
			}
		}
		nodes[v].r = dfn - 1
	}
	dfs(rt, -1)

	ts := []*node93{build93(0, n-1)}
	dis := make([]int, n)
	dis[rt] = 1
	q := []int{rt}
	for len(q) > 0 {
		tmp := q
		q = nil
		t := ts[len(ts)-1]
		for _, v := range tmp {
			t = t.update(0, n-1, nodes[v].l, a[v])
			for _, w := range g[v] {
				if dis[w] == 0 {
					dis[w] = len(ts) + 1
					q = append(q, w)
				}
			}
		}
		ts = append(ts, t)
	}

	Fscan(in, &m)
	for range m {
		Fscan(in, &x, &k)
		x = (x + ans) % n
		k = (k + ans) % n
		p := nodes[x]
		ans = ts[min(dis[x]+k, len(ts)-1)].query(0, n-1, p.l, p.r)
		Fprintln(out, ans)
	}
}

//func main() { debug.SetGCPercent(-1); cf893F(bufio.NewReader(os.Stdin), os.Stdout) }
