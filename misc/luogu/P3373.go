package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
const mod3373 int = 571373

type seg3373 []struct{ l, r, mulTodo, addTodo, s int }

func (t seg3373) maintain(o int) {
	t[o].s = (t[o<<1].s + t[o<<1|1].s) % mod3373
}

func (t seg3373) build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].mulTodo = l, r, 1
	if l == r {
		t[o].s = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg3373) doMul(o, v int) {
	to := &t[o]
	to.mulTodo = to.mulTodo * v % mod3373
	to.addTodo = to.addTodo * v % mod3373
	to.s = to.s * v % mod3373
}
func (t seg3373) doAdd(o, v int) {
	to := &t[o]
	to.addTodo = (to.addTodo + v) % mod3373
	to.s = (to.s + (to.r-to.l+1)*v) % mod3373
}
func (t seg3373) spread(o int) {
	if v := t[o].mulTodo; v != 1 {
		t.doMul(o<<1, v)
		t.doMul(o<<1|1, v)
		t[o].mulTodo = 1
	}
	if v := t[o].addTodo; v != 0 {
		t.doAdd(o<<1, v)
		t.doAdd(o<<1|1, v)
		t[o].addTodo = 0
	}
}

func (t seg3373) update(o, l, r, v int, mul bool) {
	if l <= t[o].l && t[o].r <= r {
		if mul {
			t.doMul(o, v)
		} else {
			t.doAdd(o, v)
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v, mul)
	}
	if m < r {
		t.update(o<<1|1, l, r, v, mul)
	}
	t.maintain(o)
}

func (t seg3373) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].s
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return (t.query(o<<1, l, r) + t.query(o<<1|1, l, r)) % mod3373
}

func p3373(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var ignore, n, q, op, l, r, v int
	Fscan(in, &n, &q, &ignore)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg3373, 4*n)
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		if Fscan(in, &op, &l, &r); op < 3 {
			Fscan(in, &v)
			t.update(1, l, r, v, op == 1)
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { p3373(os.Stdin, os.Stdout) }
