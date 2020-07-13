package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct {
	l, r               int
	ans, sum, pre, suf int64
}

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.init(a)
	return t
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (t seg) _set(o, val int) {
	v := int64(val)
	t[o].sum = v
	v = max(v, 0)
	t[o].ans = v
	t[o].pre = v
	t[o].suf = v
}

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].ans = max(max(lo.ans, ro.ans), lo.suf+ro.pre)
	t[o].sum = lo.sum + ro.sum
	t[o].pre = max(lo.pre, lo.sum+ro.pre)
	t[o].suf = max(ro.suf, ro.sum+lo.suf)
}

func (t seg) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t._set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, idx, val int) {
	if t[o].l == t[o].r {
		t._set(o, val)
		return
	}
	if idx <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, idx, val)
	} else {
		t._update(o<<1|1, idx, val)
	}
	t._pushUp(o)
}

func (t seg) init(a []int)        { t._build(a, 1, 1, len(a)) }
func (t seg) update(idx, val int) { t._update(1, idx, val) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, i, v int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTree(a)
	Fprintln(out, t[1].ans)
	for ; q > 0; q-- {
		Fscan(in, &i, &v)
		t.update(i+1, v)
		Fprintln(out, t[1].ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
