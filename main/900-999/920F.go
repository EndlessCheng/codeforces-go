package main

import (
	"bufio"
	. "fmt"
	"io"
)

var d920 [1e6 + 1]int

type seg920 []struct {
	l, r  int
	s     int64
	fixed bool
}

func (t seg920) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].s = lo.s + ro.s
	t[o].fixed = lo.fixed && ro.fixed
}

func (t seg920) _build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].s = int64(a[l-1])
		t[o].fixed = a[l-1] <= 2
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg920) _update(o, l, r int) {
	to := &t[o]
	if to.fixed {
		return
	}
	if to.l == to.r {
		to.s = int64(d920[to.s])
		to.fixed = to.s <= 2
		return
	}
	m := (to.l + to.r) >> 1
	if l <= m {
		t._update(o<<1, l, r)
	}
	if m < r {
		t._update(o<<1|1, l, r)
	}
	t._pushUp(o)
}

func (t seg920) _query(o, l, r int) (s int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		return t[o].s
	}
	m := (ol + or) >> 1
	if l <= m {
		s += t._query(o<<1, l, r)
	}
	if m < r {
		s += t._query(o<<1|1, l, r)
	}
	return
}

func (t seg920) init(a []int)         { t._build(a, 1, 1, len(a)) }
func (t seg920) update(l, r int)      { t._update(1, l, r) }
func (t seg920) query(l, r int) int64 { return t._query(1, l, r) }

// github.com/EndlessCheng/codeforces-go
func CF920F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx = 1e6
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			d920[j]++
		}
	}

	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg920, 4*n)
	t.init(a)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			t.update(l, r)
		} else {
			Fprintln(out, t.query(l, r))
		}
	}
}

//func main() { CF920F(os.Stdin, os.Stdout) }
