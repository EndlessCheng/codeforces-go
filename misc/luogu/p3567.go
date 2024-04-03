package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
type seg3567 []struct {
	l, r int
	v, c int
}

func (t seg3567) do(v, c, w, d int) (int, int) {
	if v == w {
		return v, c + d
	}
	if c > d {
		return v, c - d
	}
	return w, d - c
}

func (t seg3567) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].v, t[o].c = t.do(lo.v, lo.c, ro.v, ro.c)
}

func (t seg3567) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].v, t[o].c = a[l-1], 1
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg3567) query(o, l, r int) (int, int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].v, t[o].c
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	a, b := t.query(o<<1, l, r)
	c, d := t.query(o<<1|1, l, r)
	return t.do(a, b, c, d)
}

func p3567(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, buf := 1<<12, make([]byte, 1<<12)
	rc := func() byte {
		if _i == 1<<12 {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, q := ri(), ri()
	a := make([]int, n)
	ps := make([][]int, n+1)
	for i := range a {
		a[i] = ri()
		ps[a[i]] = append(ps[a[i]], i+1)
	}
	t := make(seg3567, 4*len(a))
	t.build(a, 1, 1, len(a))
	for ; q > 0; q-- {
		l, r := ri(), ri()
		v, _ := t.query(1, l, r)
		if sort.SearchInts(ps[v][sort.SearchInts(ps[v], l):], r+1) > (r-l+1)/2 {
			Fprintln(out, v)
		} else {
			Fprintln(out, 0)
		}
	}
}

//func main() { p3567(os.Stdin, os.Stdout) }
