package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, min, cnt int }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	if lo.min == ro.min {
		t[o].min, t[o].cnt = lo.min, lo.cnt+ro.cnt
	} else if lo.min < ro.min {
		t[o].min, t[o].cnt = lo.min, lo.cnt
	} else {
		t[o].min, t[o].cnt = ro.min, ro.cnt
	}
}

func (t seg) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = a[l-1]
		t[o].cnt = 1
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, idx, val int) {
	if t[o].l == t[o].r {
		t[o].min = val
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) (int, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min, t[o].cnt
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	minL, cntL := t._query(o<<1, l, r)
	minR, cntR := t._query(o<<1|1, l, r)
	if minL == minR {
		return minL, cntL + cntR
	} else if minL < minR {
		return minL, cntL
	} else {
		return minR, cntR
	}
}

func (t seg) init(a []int)              { t._build(a, 1, 1, len(a)) }
func (t seg) update(idx, val int)       { t._update(1, idx, val) }
func (t seg) query(l, r int) (int, int) { return t._query(1, l, r) }

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
			v, c := t.query(l+1, r)
			Fprintln(out, v, c)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
