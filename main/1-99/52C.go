package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg52 []struct {
	l, r      int
	min, todo int64
}

func (t seg52) maintain(o int) {
	t[o].min = min52(t[o<<1].min, t[o<<1|1].min)
}

func (t seg52) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = int64(a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg52) spread(o int) {
	if v := t[o].todo; v != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.min += v
		lo.todo += v
		ro.min += v
		ro.todo += v
		t[o].todo = 0
	}
}

func (t seg52) update(o, l, r, v int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].min += int64(v)
		t[o].todo += int64(v)
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

func (t seg52) query(o, l, r int) int64 {
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
	return min52(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF52C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r, v int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg52, 4*n)
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
				Fprintln(out, min52(t.query(1, l, n), t.query(1, 1, r)))
			} else {
				Fprintln(out, t.query(1, l, r))
			}
		}
	}
}

func min52(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//func main() { CF52C(os.Stdin, os.Stdout) }


