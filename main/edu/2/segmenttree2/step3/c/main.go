package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct{ l, r, todo, max int }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) { t[o].max = max(t[o<<1].max, t[o<<1|1].max) }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(o, v int) {
	to := &t[o]
	to.todo += v
	to.max += v
}

func (t seg) spread(o int) {
	if v := t[o].todo; v > 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
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

func (t seg) query(o, l, x int) int {
	if x > t[o].max {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	if l <= t[o<<1].r && x <= t[o<<1].max {
		if p := t.query(o<<1, l, x); p > 0 {
			return p
		}
	}
	return t.query(o<<1|1, l, x)
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
		if Fscan(in, &op); op == 1 {
			Fscan(in, &l, &r, &v)
			t.update(1, l+1, r, v)
		} else {
			Fscan(in, &v, &l)
			Fprintln(out, t.query(1, l+1, v)-1)
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
