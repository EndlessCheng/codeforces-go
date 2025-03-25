package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct{ mx1, mx2, c1, c2 int }

type seg []struct {
	l, r int
	val  data
}

func mergeInfo(a, b data) data {
	if a.mx1 < b.mx1 {
		a, b = b, a
	}
	if a.mx1 > b.mx1 {
		if b.mx1 > a.mx2 {
			a.mx2 = b.mx1
			a.c2 = b.c1
		} else if b.mx1 == a.mx2 {
			a.c2 += b.c1
		}
	} else {
		a.c1 += b.c1
		if b.mx2 > a.mx2 {
			a.mx2 = b.mx2
			a.c2 = b.c2
		} else if b.mx2 == a.mx2 {
			a.c2 += b.c2
		}
	}
	return a
}

func (t seg) maintain(o int) {
	t[o].val = mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val.mx1 = a[l]
		t[o].val.c1 = 1
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].val.mx1 = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) data {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func newSegmentTreeWithArray(a []int) seg {
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
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTreeWithArray(a)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			t.update(1, l-1, r)
		} else {
			Fprintln(out, t.query(1, l-1, r-1).c2)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
