package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
const inf int = 1e18

type mat [5][5]int

type seg []struct {
	l, r int
	val  mat
}

func newVal(a, b int) mat {
	return mat{
		{0, a + b, a + b*2, -inf, -inf},
		{-inf, a, a + b, -inf, -inf},
		{-inf, -inf, 0, a + b, a + b*2},
		{-inf, -inf, -inf, a, a + b},
		{-inf, -inf, -inf, -inf, 0},
	}
}

func mergeInfo(a, b mat) (c mat) {
	for i := range 5 {
		for j := range 5 {
			c[i][j] = -inf
		}
	}
	for i := range 5 {
		for k := i; k < 5; k++ {
			for j := k; j < 5; j++ {
				c[i][j] = max(c[i][j], a[i][k]+b[k][j])
			}
		}
	}
	return
}

func (t seg) build(a [][2]int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = newVal(a[l][0], a[l][1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, a, b int) {
	if t[o].l == t[o].r {
		t[o].val = newVal(a, b)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, a, b)
	} else {
		t.update(o<<1|1, i, a, b)
	}
	t.maintain(o)
}

func (t seg) maintain(o int) {
	t[o].val = mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t seg) query(o, l, r int) mat {
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

func cf2042F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, p, x int
	Fscan(in, &n)
	a := make([][2]int, n)
	for i := range a {
		Fscan(in, &a[i][0])
	}
	for i := range a {
		Fscan(in, &a[i][1])
	}

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &p, &x)
		p--
		if op < 3 {
			a[p][op-1] = x
			t.update(1, p, a[p][0], a[p][1])
		} else {
			Fprintln(out, t.query(1, p, x-1)[0][4])
		}
	}
}

//func main() { cf2042F(bufio.NewReader(os.Stdin), os.Stdout) }
