package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r    int
	todo, s int64
}

func (t seg) maintain(o int) {
	t[o].s = t[o<<1].s + t[o<<1|1].s
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o int, v int64) {
	to := &t[o]
	to.todo += v
	to.s += int64(to.r-to.l+1) * v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg) update(o, l, r int, v int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
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
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r int
	var a0, d int64
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l); op == 1 {
			Fscan(in, &r, &a0, &d)
			t.update(1, l, l, a0)
			if r > l {
				t.update(1, l+1, r, d)
			}
			if r < n {
				t.update(1, r+1, r+1, -a0-d*int64(r-l))
			}
		} else {
			Fprintln(out, t.query(1, 1, l))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
