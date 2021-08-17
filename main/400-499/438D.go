package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg38 []struct {
	l, r  int
	max   int
	sum   int64
}

func (t seg38) set(o, v int) {
	t[o].max = v
	t[o].sum = int64(v)
}

func (t seg38) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].max = max38(lo.max, ro.max)
	t[o].sum = lo.sum + ro.sum
}

func (t seg38) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg38) mod(o, l, r, mod int) {
	if t[o].max < mod {
		return
	}
	if t[o].l == t[o].r {
		t[o].max %= mod
		t[o].sum = int64(t[o].max)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.mod(o<<1, l, r, mod)
	}
	if r > m {
		t.mod(o<<1|1, l, r, mod)
	}
	t.maintain(o)
}

func (t seg38) update(o, i, v int) {
	if t[o].l == t[o].r {
		t.set(o, v)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t seg38) query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func CF438D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, mod int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg38, n*4)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			Fprintln(out, t.query(1, l, r))
		} else if op == 2 {
			Fscan(in, &mod)
			t.mod(1, l, r, mod)
		} else {
			t.update(1, l, r)
		}
	}
}

func max38(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func main() { CF438D(os.Stdin, os.Stdout) }
