package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg145 []struct {
	l, r, s int
	fixed   bool
}

func (t seg145) maintain(o int) {
	t[o].s = t[o<<1].s + t[o<<1|1].s
	t[o].fixed = t[o<<1].fixed && t[o<<1|1].fixed
}

func (t seg145) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].s = a[l]
		t[o].fixed = a[l] == 1
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg145) update(o, l, r int) {
	to := &t[o]
	if to.fixed {
		return
	}
	if to.l == to.r {
		to.s = int(math.Sqrt(float64(to.s)))
		to.fixed = to.s == 1
		return
	}
	m := (to.l + to.r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg145) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].s
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

func p4145(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, l, r int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg145, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &op, &l, &r)
		if l > r {
			l, r = r, l
		}
		if op == 0 {
			t.update(1, l-1, r-1)
		} else {
			Fprintln(out, t.query(1, l-1, r-1))
		}
	}
}

//func main() { p4145(bufio.NewReader(os.Stdin), os.Stdout) }
