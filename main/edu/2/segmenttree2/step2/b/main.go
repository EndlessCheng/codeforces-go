package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
const mod int64 = 1e9 + 7

type seg []struct {
	l, r    int
	s, todo int64
}

func (t seg) maintain(o int) {
	t[o].s = (t[o<<1].s + t[o<<1|1].s) % mod
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, 1
	if l == r {
		t[o].s = 1
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 1 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.s = lo.s * v % mod
		lo.todo = lo.todo * v % mod
		ro.s = ro.s * v % mod
		ro.todo = ro.todo * v % mod
		t[o].todo = 1
	}
}

func (t seg) update(o, l, r, v int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].s = t[o].s * int64(v) % mod
		t[o].todo = t[o].todo * int64(v) % mod
		return
	}
	t.spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].s
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return (t.query(o<<1, l, r) + t.query(o<<1|1, l, r)) % mod
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			Fscan(in, &v)
			t.update(1, l+1, r, v)
		} else {
			Fprintln(out, t.query(1, l+1, r))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
