package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, max int }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) _pushUp(o int) { t[o].max = max(t[o<<1].max, t[o<<1|1].max) }

func (t seg) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].max = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].max = val
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, x int) int {
	if t[o].max < x {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	if l <= t[o<<1].r && t[o<<1].max >= x {
		if res := t._query(o<<1, l, x); res > 0 {
			return res
		}
	}
	return t._query(o<<1|1, l, x)
}

func (t seg) init(a []int)        { t._build(a, 1, 1, len(a)) }
func (t seg) update(idx, val int) { t._update(1, idx, val) }
func (t seg) query(l, x int) int  { return t._query(1, l, x) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, v, x, l int
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
			Fscan(in, &x, &l)
			Fprintln(out, t.query(l+1, x)-1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
