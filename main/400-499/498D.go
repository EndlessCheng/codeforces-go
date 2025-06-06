package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg98 []struct {
	l, r int
	val  [60]uint32
}

func (seg98) merge(a, b [60]uint32) (c [60]uint32) {
	for i, t := range a {
		c[i] = t + b[(uint32(i)+t)%60]
	}
	return
}

func (seg98) set(v int) (c [60]uint32) {
	for i := range c {
		if i%v == 0 {
			c[i] = 2
		} else {
			c[i] = 1
		}
	}
	return c
}

func (t seg98) maintain(o int) {
	t[o].val = t.merge(t[o<<1].val, t[o<<1|1].val)
}

func (t seg98) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = t.set(a[l])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg98) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].val = t.set(val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg98) query(o, l, r int) [60]uint32 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf498D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, l, r int
	var op string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg98, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)

	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == "A" {
			Fprintln(out, t.query(1, l-1, r-2)[0])
		} else {
			t.update(1, l-1, r)
		}
	}
}

//func main() { cf498D(bufio.NewReader(os.Stdin), os.Stdout) }
