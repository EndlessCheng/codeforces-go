package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct{ inv, c0, c1 int }

type seg []struct {
	l, r int
	d    data
	todo int
}

func mergeInfo(a, b data) data {
	return data{a.inv + b.inv + a.c1*b.c0, a.c0 + b.c0, a.c1 + b.c1}
}

func (t seg) apply(o int) {
	cur := &t[o]
	d := &cur.d

	c0 := d.c0
	c1 := d.c1
	sz := c0 + c1
	d.inv = sz*(sz-1)/2 - c0*(c0-1)/2 - c1*(c1-1)/2 - d.inv
	d.c0 = c1
	d.c1 = c0

	cur.todo ^= 1
}

func (t seg) maintain(o int) {
	t[o].d = mergeInfo(t[o<<1].d, t[o<<1|1].d)
}

func (t seg) spread(o int) {
	if t[o].todo == 0 {
		return
	}
	t.apply(o << 1)
	t.apply(o<<1 | 1)
	t[o].todo = 0
}

func (t seg) build(a []data, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].d = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].d
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func newLazySegmentTreeWithArray(a []data) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]data, n)
	for i := range a {
		Fscan(in, &a[i].c1)
		a[i].c0 = a[i].c1 ^ 1
	}
	t := newLazySegmentTreeWithArray(a)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		l--
		r--
		if op == 1 {
			t.update(1, l, r)
		} else {
			Fprintln(out, t.query(1, l, r).inv)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
