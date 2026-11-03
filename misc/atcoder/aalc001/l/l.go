package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type pair struct{ s, maxS int }
type seg []struct {
	data [3]pair
	mp   [3]uint8
}

func (t seg) apply(o int, f [3]uint8) {
	p := &t[o]
	p.data = [3]pair{p.data[f[0]], p.data[f[1]], p.data[f[2]]}
	p.mp = [3]uint8{p.mp[f[0]], p.mp[f[1]], p.mp[f[2]]}
}

var mp0 = [3]uint8{0, 1, 2}

func (t seg) spread(o int) {
	f := t[o].mp
	if f == mp0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].mp = mp0
}

func (t seg) maintain(o int) {
	p := &t[o]
	lo, ro := &t[o<<1], &t[o<<1|1]
	for i := range 3 {
		p.data[i].s = lo.data[i].s + ro.data[i].s
		p.data[i].maxS = max(lo.data[i].maxS, lo.data[i].s+ro.data[i].maxS)
	}
}

func (t seg) build(a [][3]int, o, l, r int) {
	t[o].mp = mp0
	if l == r {
		for i, x := range a[l-1] {
			t[o].data[i] = pair{x, x}
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, ql, qr int, f [3]uint8) {
	if ql <= l && r <= qr {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, f)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, f)
	}
	t.maintain(o)
}

var s int

func (t seg) findFirst(o, l, r, ql int) int {
	if l == r {
		s += t[o].data[0].s
		if s > 0 {
			return l
		}
		return 0
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		i := t.findFirst(o<<1, l, m, ql)
		if i > 0 {
			return i
		}
		p := t[o<<1|1].data[0]
		if s+p.maxS <= 0 {
			s += p.s
			return 0
		}
	}
	return t.findFirst(o<<1|1, m+1, r, ql)
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([][3]int, n)
	for j := range 3 {
		for i := range a {
			Fscan(in, &a[i][j])
		}
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)

	mps := [3][3]uint8{{1, 0, 2}, {0, 2, 1}, {2, 1, 0}}
	for range q {
		Fscan(in, &op, &l)
		if op < 4 {
			Fscan(in, &r)
			t.update(1, 1, n, l, r, mps[op-1])
		} else {
			s = 0
			i := t.findFirst(1, 1, n, l)
			if i == 0 {
				i = n + 1
			}
			Fprintln(out, i)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
