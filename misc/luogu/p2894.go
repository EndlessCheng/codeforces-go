package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg2894 []struct{ l, r, pre0, suf0, max0, todo int }

func newSegTree(n int) seg2894 {
	t := make(seg2894, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	return t
}

func (t seg2894) do(i, v int) {
	o := &t[i]
	size := 0
	if v <= 0 {
		size = o.r - o.l + 1
	}
	o.pre0 = size
	o.suf0 = size
	o.max0 = size
	o.todo = v
}

func (t seg2894) spread(o int) {
	v := t[o].todo
	if v != -1 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = -1
	}
}

func (t seg2894) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t.do(o, -1)
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg2894) update(o, l, r, v int) {
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

	lo, ro := t[o<<1], t[o<<1|1]
	t[o].pre0 = lo.pre0
	if lo.pre0 == m-t[o].l+1 {
		t[o].pre0 += ro.pre0
	}
	t[o].suf0 = ro.suf0
	if ro.suf0 == t[o].r-m {
		t[o].suf0 += lo.suf0
	}
	t[o].max0 = max(lo.max0, ro.max0, lo.suf0+ro.pre0)
}

func (t seg2894) findFirst(o, size int) int {
	if t[o].max0 < size {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	idx := t.findFirst(o<<1, size)
	if idx == 0 {
		if t[o<<1].suf0+t[o<<1|1].pre0 >= size {
			m := (t[o].l + t[o].r) >> 1
			return m - t[o<<1].suf0 + 1
		}
		idx = t.findFirst(o<<1|1, size)
	}
	return idx
}

func p2894(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, x, c int
	Fscan(in, &n, &m)
	t := make(seg2894, 2<<bits.Len(uint(n)))
	t.build(1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &op, &x)
		if op == 1 {
			i := t.findFirst(1, x)
			if i > 0 {
				t.update(1, i, i+x-1, 1)
			}
			Fprintln(out, i)
		} else {
			Fscan(in, &c)
			t.update(1, x, x+c-1, 0)
		}
	}
}

//func main() { p2894(bufio.NewReader(os.Stdin), os.Stdout) }
