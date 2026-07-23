package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct{ min, cntNZ int }
type seg []struct {
	data
	dec int
}

func (t seg) maintain(o int) {
	l, r := t[o<<1].data, t[o<<1|1].data
	t[o].min = min(l.min, r.min)
	t[o].cntNZ = l.cntNZ + r.cntNZ
}

func (t seg) apply(o, f int) {
	t[o].min -= f
	t[o].dec += f
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
		t[o].cntNZ = 1
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, ql, qr, k int) (buy int) {
	if t[o].min == 1e18 {
		return
	}
	if t[o].min > k && ql <= l && r <= qr {
		t.apply(o, k)
		return t[o].cntNZ * k
	}
	if l == r {
		buy = t[o].min
		t[o].min = 1e18
		t[o].cntNZ = 0
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
