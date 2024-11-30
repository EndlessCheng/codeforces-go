package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type data struct{ ans, sum, pre, suf int }
type seg []struct {
	l, r int
	data
}

func (t seg) set(o, v int) { t[o].data = data{v, v, v, v} }

func (t seg) do(lo, ro data) (o data) {
	o.ans = max(lo.ans, ro.ans, lo.suf+ro.pre)
	o.sum = lo.sum + ro.sum
	o.pre = max(lo.pre, lo.sum+ro.pre)
	o.suf = max(ro.suf, ro.sum+lo.suf)
	return
}

func (t seg) maintain(o int) { t[o].data = t.do(t[o<<1].data, t[o<<1|1].data) }

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) (d data) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.do(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p4513(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg, 4*n)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			if l > r {
				l, r = r, l
			}
			Fprintln(out, t.query(1, l, r).ans)
		} else {
			t.update(1, l, r)
		}
	}
}

//func main() { p4513(os.Stdin, os.Stdout) }
