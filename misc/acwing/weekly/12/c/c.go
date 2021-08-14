package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct{ l, r, min, todo int }

func (t seg) maintain(o int) {
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) spread(o int) {
	if v := t[o].todo; v != 0 {
		t[o<<1].min += v
		t[o<<1].todo += v
		t[o<<1|1].min += v
		t[o<<1|1].todo += v
		t[o].todo = 0
	}
}

func (t seg) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t[o].min += v
		t[o].todo += v
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

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg, n*4)
	t.build(a, 1, 1, n)
	for Fscanf(in, "\n%d\n", &q); q > 0; q-- {
		m, _ := Fscanln(in, &l, &r, &v)
		l++
		r++
		if m == 3 {
			if l > r {
				t.update(1, l, n, v)
				t.update(1, 1, r, v)
			} else {
				t.update(1, l, r, v)
			}
		} else {
			if l > r {
				Fprintln(out, min(t.query(1, l, n), t.query(1, 1, r)))
			} else {
				Fprintln(out, t.query(1, l, r))
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
