package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
const mod3373 int = 571373

type seg3373 []struct{ l, r, sum, mulTodo, addTodo int }

func newLazySegmentTree(a []int) seg3373 {
	t := make(seg3373, 4*len(a))
	t.init(a)
	return t
}

func (t seg3373) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = (lo.sum + ro.sum) % mod3373
}

func (t seg3373) _build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].mulTodo = l, r, 1
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg3373) _spread(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	if mul := t[o].mulTodo; mul != 1 {
		lo.sum = lo.sum * mul % mod3373
		ro.sum = ro.sum * mul % mod3373
		lo.mulTodo = lo.mulTodo * mul % mod3373
		ro.mulTodo = ro.mulTodo * mul % mod3373
		lo.addTodo = lo.addTodo * mul % mod3373
		ro.addTodo = ro.addTodo * mul % mod3373
		t[o].mulTodo = 1
	}
	if add := t[o].addTodo; add != 0 {
		lo.sum = (lo.sum + add*(lo.r-lo.l+1)) % mod3373
		ro.sum = (ro.sum + add*(ro.r-ro.l+1)) % mod3373
		lo.addTodo = (lo.addTodo + add) % mod3373
		ro.addTodo = (ro.addTodo + add) % mod3373
		t[o].addTodo = 0
	}
}

func (t seg3373) _updateMul(o, l, r, mul int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum = t[o].sum * mul % mod3373
		t[o].mulTodo = t[o].mulTodo * mul % mod3373
		t[o].addTodo = t[o].addTodo * mul % mod3373
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._updateMul(o<<1, l, r, mul)
	}
	if m < r {
		t._updateMul(o<<1|1, l, r, mul)
	}
	t._pushUp(o)
}

func (t seg3373) _updateAdd(o, l, r, add int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum = (t[o].sum + add*(or-ol+1)) % mod3373
		t[o].addTodo = (t[o].addTodo + add) % mod3373
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._updateAdd(o<<1, l, r, add)
	}
	if m < r {
		t._updateAdd(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t seg3373) _query(o, l, r int) (res int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		return t[o].sum
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		res += t._query(o<<1, l, r)
	}
	if m < r {
		res += t._query(o<<1|1, l, r)
	}
	res %= mod3373
	return
}

func (t seg3373) init(a []int)            { t._build(a, 1, 1, len(a)) }  // starts at 0
func (t seg3373) updateMul(l, r, val int) { t._updateMul(1, l, r, val) } // [l,r] 1<=l<=r<=n
func (t seg3373) updateAdd(l, r, val int) { t._updateAdd(1, l, r, val) } // [l,r] 1<=l<=r<=n
func (t seg3373) query(l, r int) int      { return t._query(1, l, r) }   // [l,r] 1<=l<=r<=n

func p3373(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var ignore, n, q, op, l, r, val int
	Fscan(in, &n, &q, &ignore)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := newLazySegmentTree(a)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fscan(in, &val)
			t.updateMul(l, r, val)
		} else if op == 2 {
			Fscan(in, &val)
			t.updateAdd(l, r, val)
		} else {
			Fprintln(out, t.query(l, r))
		}
	}
}

//func main() { run(os.Stdin, os.Stdout) }
