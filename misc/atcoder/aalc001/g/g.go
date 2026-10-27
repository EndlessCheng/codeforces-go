package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
const mod = 998244353

type seg []struct {
	s   [4]int
	add int
}

func (t seg) apply(o, sz, x int) {
	p := &t[o]
	x2 := x * x % mod
	x3 := x2 * x % mod
	x4 := x3 * x % mod
	p.s[3] = (p.s[3] + (p.s[2]*x+p.s[0]*x3)*4%mod + p.s[1]*x2*6 + sz*x4) % mod
	p.s[2] = (p.s[2] + p.s[1]*x*3 + p.s[0]*x2*3 + sz*x3) % mod
	p.s[1] = (p.s[1] + p.s[0]*x*2 + sz*x2) % mod
	p.s[0] = (p.s[0] + x*sz) % mod
	p.add = (p.add + x) % mod
}

func (t seg) spread(o, l, r int) {
	f := t[o].add
	if f == 0 {
		return
	}
	m := (l + r) >> 1
	t.apply(o<<1, m-l+1, f)
	t.apply(o<<1|1, r-m, f)
	t[o].add = 0
}

func (t seg) maintain(o int) {
	for i := range 4 {
		t[o].s[i] = (t[o<<1].s[i] + t[o<<1|1].s[i]) % mod
	}
}

func (t seg) build(in io.Reader, o, l, r int) {
	if l == r {
		var x int
		Fscan(in, &x)
		t[o].s[0] = x
		t[o].s[1] = x * x % mod
		t[o].s[2] = t[o].s[1] * x % mod
		t[o].s[3] = t[o].s[2] * x % mod
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, ql, qr, x int) {
	if ql <= l && r <= qr {
		t.apply(o, r-l+1, x)
		return
	}
	t.spread(o, l, r)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, x)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, x)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r, ql, qr, k int) int {
	if ql <= l && r <= qr {
		return t[o].s[k]
	}
	t.spread(o, l, r)
	m := (l + r) >> 1
	if qr <= m {
		return t.query(o<<1, l, m, ql, qr, k)
	}
	if ql > m {
		return t.query(o<<1|1, m+1, r, ql, qr, k)
	}
	return t.query(o<<1, l, m, ql, qr, k) + t.query(o<<1|1, m+1, r, ql, qr, k)
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, x int
	Fscan(in, &n, &q)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)
	for range q {
		Fscan(in, &op, &l, &r, &x)
		if op == 1 {
			t.update(1, 1, n, l, r, x)
		} else {
			Fprintln(out, t.query(1, 1, n, l, r, x-1)%mod)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
