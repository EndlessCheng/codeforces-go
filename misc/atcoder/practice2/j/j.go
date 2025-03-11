package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type seg []struct {
	l, r int
	val  int
}

func mergeInfo(a, b int) int {
	return max(a, b)
}

func setInfo(a, b int) int {
	return b
}

func (t seg) maintain(o int) {
	t[o].val = mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int, val int) {
	if t[o].l == t[o].r {
		t[o].val = setInfo(t[o].val, val)
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

func (t seg) query(o, l, r int) int {
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
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func (t seg) findFirst(o, l, r int, f func(int) bool) int {
	if t[o].l > r || t[o].r < l || !f(t[o].val) {
		return -1
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	idx := t.findFirst(o<<1, l, r, f)
	if idx < 0 {
		idx = t.findFirst(o<<1|1, l, r, f)
	}
	return idx
}

func newSegmentTreeWithArray(a []int) seg {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	return t
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newSegmentTreeWithArray(a)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			t.update(1, l-1, r)
		} else if op == 2 {
			Fprintln(out, t.query(1, l-1, r-1))
		} else {
			i := t.findFirst(1, l-1, n-1, func(nodeMax int) bool { return nodeMax >= r })
			if i < 0 {
				i = n
			}
			Fprintln(out, i+1)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
