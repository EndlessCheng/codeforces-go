package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg10 []struct{ l, r, min, todo int }

func (t seg10) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg10) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg10) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg10) update(o, l, r, v int) {
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

func (t seg10) maintain(o int) {
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg10) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1110F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, dfn int
	Fscan(in, &n, &q)
	type pair struct{ to, wt int }
	g := make([][]pair, n)
	for w := 1; w < n; w++ {
		var v, wt int
		Fscan(in, &v, &wt)
		g[v-1] = append(g[v-1], pair{w, wt})
	}
	a := make([]int, n)
	nodes := make([]struct{ l, r int }, n)
	var build func(int, int) int
	build = func(v, d int) (size int) {
		if g[v] == nil {
			a[dfn] = d
		} else {
			a[dfn] = 1e18
		}
		dfn++
		nodes[v].l = dfn
		for _, e := range g[v] {
			size += build(e.to, d+e.wt)
		}
		nodes[v].r = nodes[v].l + size
		size++
		return
	}
	build(0, 0)

	t := make(seg10, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	type query struct{ l, r, i int }
	qs := make([][]query, n)
	for i := 0; i < q; i++ {
		var v, l, r int
		Fscan(in, &v, &l, &r)
		qs[v-1] = append(qs[v-1], query{l, r, i})
	}
	ans := make([]int, q)
	var f func(int)
	f = func(v int) {
		for _, q := range qs[v] {
			ans[q.i] = t.query(1, q.l, q.r)
		}
		for _, e := range g[v] {
			p := nodes[e.to]
			t.update(1, 1, n, e.wt)
			t.update(1, p.l, p.r, -e.wt*2)
			f(e.to)
			t.update(1, 1, n, -e.wt)
			t.update(1, p.l, p.r, e.wt*2)
		}
	}
	f(0)
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1110F(os.Stdin, os.Stdout) }
