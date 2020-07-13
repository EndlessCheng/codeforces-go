package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct {
	l, r int
	val  int64
}

func newSegmentTree(a []int64) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func (t seg) _pushUp(o int) { t[o].val = t[o<<1].val + t[o<<1|1].val }

func (t seg) _build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, idx int, val int64) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	return t._query(o<<1, l, r) + t._query(o<<1|1, l, r)
}

func (t seg) init(a []int64)            { t._build(a, 1, 1, len(a)) }
func (t seg) update(idx int, val int64) { t._update(1, idx, val) }
func (t seg) query(l, r int) int64      { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, v, l, r int
	Fscan(in, &n, &q)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTree(a)
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &i, &v)
			t.update(i+1, int64(v))
		} else {
			Fscan(in, &l, &r)
			Fprintln(out, t.query(l+1, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
