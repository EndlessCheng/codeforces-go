package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, v, todo int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) spread(o int) {
	if v := t[o].todo; v >= 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.v = v
		lo.todo = v
		ro.v = v
		ro.todo = v
		t[o].todo = -1
	}
}

func (t seg) update(o, l, r, v int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].v = v
		t[o].todo = v
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
}

func (t seg) query(o, i int) int {
	if t[o].l == t[o].r {
		return t[o].v
	}
	t.spread(o)
	if i <= (t[o].l+t[o].r)>>1 {
		return t.query(o<<1, i)
	}
	return t.query(o<<1|1, i)
}

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, v int
	Fscan(in, &n, &q)
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l); op == 1 {
			Fscan(in, &r, &v)
			t.update(1, l+1, r, v)
		} else {
			Fprintln(out, t.query(1, l+1))
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
