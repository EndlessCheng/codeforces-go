package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct{ l, r, max, todo int }

func (t seg) maintain(o int) {
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
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

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.max += v
		lo.todo += v
		ro.max += v
		ro.todo += v
		t[o].todo = 0
	}
}

func (t seg) update(o, l, r, v int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].max += v
		t[o].todo += v
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

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, l, r, v int
	Fscan(in, &n, &m)
	type pair struct{ p, v int }
	add := make([][]pair, n+1)
	for ; m > 0; m-- {
		Fscan(in, &l, &r, &v)
		add[r] = append(add[r], pair{l, v})
	}

	t := make(seg, 4*n)
	t.build(1, 1, n)
	for i := 1; i <= n; i++ {
		mx := t.query(1, 1, i)
		t.update(1, i, i, mx)
		for _, p := range add[i] {
			t.update(1, p.p, i, p.v)
		}
	}
	Fprint(out, max(0, t[1].max))
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
