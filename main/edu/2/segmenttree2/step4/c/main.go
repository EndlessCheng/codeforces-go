package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct {
	l, r, todo, c, s int
	lb, rb           bool
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].s = lo.s + ro.s
	t[o].c = lo.c + ro.c
	if lo.rb && ro.lb {
		t[o].c--
	}
	t[o].lb = lo.lb
	t[o].rb = ro.rb
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o, v int) {
	to := &t[o]
	to.todo = v
	to.c = v
	to.s = (to.r - to.l + 1) * v
	to.lb = v > 0
	to.rb = v > 0
}

func (t seg) spread(o int) {
	if v := t[o].todo; v >= 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = -1
	}
}

func (t seg) update(o, l, r, v int) {
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

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const n int = 1e6 + 1
	t := make(seg, 4*n)
	t.build(1, 1, n)
	var q, l, d int
	var s string
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &s, &l, &d)
		l += 5e5 + 1
		r := l + d - 1
		if s[0] == 'W' {
			t.update(1, l, r, 0)
		} else {
			t.update(1, l, r, 1)
		}
		Fprintln(out, t[1].c, t[1].s)
	}
}

func main() { run(os.Stdin, os.Stdout) }
