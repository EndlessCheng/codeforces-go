package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go

type data struct{ s, ps int64 }
type seg []struct {
	l, r, todo int
	data
}

func (t seg) calc(a, b data, ln int) data {
	return data{a.s + b.s, a.ps + b.ps + int64(ln)*b.s}
}

func (t seg) maintain(o int) {
	lo := t[o<<1]
	t[o].data = t.calc(lo.data, t[o<<1|1].data, lo.r-lo.l+1)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		v := int64(a[l-1])
		t[o].s = v
		t[o].ps = v
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) do(o, v int) {
	to := &t[o]
	to.todo += v
	n := int64(to.r - to.l + 1)
	to.s += n * int64(v)
	to.ps += n * (n + 1) / 2 * int64(v)
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) update(o, l, r, v int) {
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

func (t seg) query(o, l, r int) (data, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data, t[o].r - t[o].l + 1
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	a, ln := t.query(o<<1, l, r)
	b, rn := t.query(o<<1|1, l, r)
	return t.calc(a, b, ln), ln + rn
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg, 4*n)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			Fscan(in, &v)
			t.update(1, l, r, v)
		} else {
			res, _ := t.query(1, l, r)
			Fprintln(out, res.ps)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
