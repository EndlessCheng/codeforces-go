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

type pair struct{ k, c int }

type lazySeg []struct {
	l, r int
	sum  int
	todo pair
}

func mergeInfo(a, b int) int {
	return (a + b) % mod
}

var todoInit = pair{1, 0}

func mergeTodo(f, old pair) pair {
	return pair{f.k * old.k % mod, (f.k*old.c + f.c) % mod}
}

func (t lazySeg) apply(o int, f pair) {
	cur := &t[o]

	sz := cur.r - cur.l + 1
	cur.sum = (f.k*cur.sum + f.c*sz) % mod

	cur.todo = mergeTodo(f, cur.todo)
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

func (t lazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].sum = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazySeg) update(o, l, r int, f pair) {
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

func newLazySegmentTreeWithArray(a []int) lazySeg {
	n := len(a)
	t := make(lazySeg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r, k, c int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newLazySegmentTreeWithArray(a)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		r--
		if op == 0 {
			Fscan(in, &k, &c)
			t.update(1, l, r, pair{k, c})
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
