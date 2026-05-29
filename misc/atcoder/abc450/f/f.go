package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
const mod = 998244353

type lazySeg []struct{ l, r, sum, todo int }

func mergeInfo(l, r int) int {
	return (l + r) % mod
}

const todoInit = 1

func (t lazySeg) apply(o, f int) {
	cur := &t[o]
	cur.sum = cur.sum * f % mod
	cur.todo = cur.todo * f % mod
}

func (t lazySeg) maintain(o int) {
	t[o].sum = mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t lazySeg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t lazySeg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t lazySeg) set(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].sum = v
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.set(o<<1, i, v)
	} else {
		t.set(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t lazySeg) update(o, l, r, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t lazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	type pair struct{ l, r int }
	a := make([]pair, m)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.l - b.l })

	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	t.set(1, 1, 1)
	for _, p := range a {
		t.update(1, p.r, n, 2)
		t.set(1, p.r, t.query(1, p.l, p.r))
	}
	Fprint(out, t.query(1, n, n))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
