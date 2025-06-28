package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type info77 struct{ v, i int }

type seg77 []struct {
	l, r int
	max  info77
	todo int
}

func (seg77) merge(a, b info77) info77 {
	if a.v > b.v || a.v == b.v && a.i > b.i {
		return a
	}
	return b
}

func (t seg77) apply(o, f int) {
	t[o].max.v += f
	t[o].todo += f
}

func (t seg77) maintain(o int) {
	t[o].max = t.merge(t[o<<1].max, t[o<<1|1].max)
}

func (t seg77) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg77) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].max.i = r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg77) update(o, l, r, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func cf377D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 3e5
	var n, l, v, r, ans, ll, rr int
	Fscan(in, &n)
	type worker struct{ l, v, r int }
	a := make([]worker, n)
	type tuple struct{ v, r, delta int }
	g := [mx + 2][]tuple{}
	for i := range a {
		Fscan(in, &l, &v, &r)
		a[i] = worker{l, v, r}
		g[l] = append(g[l], tuple{v, r, 1})
		g[v+1] = append(g[v+1], tuple{v, r, -1})
	}

	t := make(seg77, 2<<bits.Len(mx))
	t.build(1, 1, mx)
	for i, g := range g {
		for _, p := range g {
			t.update(1, p.v, p.r, p.delta)
		}
		if t[1].max.v > ans {
			ans = t[1].max.v
			ll, rr = i, t[1].max.i
		}
	}

	Fprintln(out, ans)
	for i, w := range a {
		if w.l <= ll && ll <= w.v && w.v <= rr && rr <= w.r {
			Fprint(out, i+1, " ")
		}
	}
}

//func main() { cf377D(bufio.NewReader(os.Stdin), os.Stdout) }
