package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
const mod18 int64 = 1e9 + 7

type matrix18 [2][2]int64

var id18 = matrix18{{1, 0}, {0, 1}}
var trans18 = matrix18{{1, 1}, {1, 0}}
var base18 = matrix18{{0, 0}, {1, 0}}

func (a matrix18) add(b matrix18) matrix18 {
	for i, r := range a {
		for j, v := range r {
			b[i][j] = (b[i][j] + v) % mod18
		}
	}
	return b
}

func (a matrix18) mul(b matrix18) (c matrix18) {
	for i, r := range a {
		for j := range b[0] {
			for k, v := range r {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod18
			}
		}
	}
	return c
}

func (a matrix18) pow(n int) matrix18 {
	res := id18
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

type seg18 []struct {
	l, r int
	todo matrix18
	sum  matrix18
}

func (t seg18) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = lo.sum.add(ro.sum)
}

func (t seg18) build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, id18
	if l == r {
		t[o].sum = trans18.pow(a[l-1]).mul(base18)
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg18) do(o int, m matrix18) {
	to := &t[o]
	to.todo = m.mul(to.todo)
	to.sum = m.mul(to.sum)
}

func (t seg18) spread(o int) {
	if m := t[o].todo; m != id18 {
		t.do(o<<1, m)
		t.do(o<<1|1, m)
		t[o].todo = id18
	}
}

func (t seg18) update(o, l, r int, mat matrix18) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, mat)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, mat)
	}
	if m < r {
		t.update(o<<1|1, l, r, mat)
	}
	t.maintain(o)
}

func (t seg18) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum[0][0]
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return (t.query(o<<1, l, r) + t.query(o<<1|1, l, r)) % mod18
}

func CF718C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, x int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg18, n*4)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op == 1 {
			Fscan(in, &x)
			t.update(1, l, r, trans18.pow(x))
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { CF718C(os.Stdin, os.Stdout) }
