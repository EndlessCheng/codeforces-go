package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct {
	min  int
	all0 bool
}
type seg []struct {
	data
	dec int
}

func (t seg) maintain(o int) {
	l, r := t[o<<1].data, t[o<<1|1].data
	t[o].min = min(l.min, r.min)
	t[o].all0 = l.all0 && r.all0
}

func (t seg) apply(o, f int) {
	cur := &t[o]
	cur.min -= f
	cur.dec += f
}

func (t seg) spread(o int) {
	f := t[o].dec
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].dec = 0
}

func (t seg) build(in io.Reader, o, l, r int) {
	if l == r {
		Fscan(in, &t[o].min)
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, ql, qr, k int) (buy int) {
	if t[o].min > k && ql <= l && r <= qr {
		t.apply(o, k)
		return (r - l + 1) * k
	}
	if l == r {
		buy = t[o].min
		t[o].min = 0
		t[o].all0 = true
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		buy += t.update(o<<1, l, m, ql, qr, k)
	}
	if m < qr {
		buy += t.update(o<<1|1, m+1, r, ql, qr, k)
	}
	t.maintain(o)
	return
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l, r, k int
	Fscan(in, &n)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)
	Fscan(in, &q)
	for range q {
		Fscan(in, &l, &r, &k)
		Fprintln(out, t.update(1, 1, n, l, r, k))
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
