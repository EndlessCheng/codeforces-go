package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, val int }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (t seg) _pushUp(o int) { t[o].val = min(t[o<<1].val, t[o<<1|1].val) }

func (t seg) _build(a []int, o, l, r int) {
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

func (t seg) _update(o, idx, val int) {
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

func (t seg) _query(o, l, r int) int {
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
	return min(t._query(o<<1, l, r), t._query(o<<1|1, l, r))
}

func (t seg) init(a []int)        { t._build(a, 1, 1, len(a)) }
func (t seg) update(idx, val int) { t._update(1, idx, val) }
func (t seg) query(l, r int) int  { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, v, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTree(a)
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &i, &v)
			t.update(i+1, v)
		} else {
			Fscan(in, &l, &r)
			Fprintln(out, t.query(l+1, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
