package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
var todoInit94 [10]int

func init() {
	for i := range todoInit94 {
		todoInit94[i] = i
	}
}

type seg94 []struct {
	l, r int
	sum  [10]int
	todo [10]int
}

func (seg94) mergeInfo(a, b [10]int) [10]int {
	for i := range a {
		a[i] += b[i]
	}
	return a
}

func (t seg94) do(O int, trans [10]int) {
	o := &t[O]
	s := [10]int{}
	for i, t := range trans {
		s[t] += o.sum[i]
	}
	o.sum = s
	for i := range trans {
		o.todo[i] = trans[o.todo[i]]
	}
}

func (t seg94) spread(o int) {
	if v := t[o].todo; v != todoInit94 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit94
	}
}

func (t seg94) build(a []int, o, l, r int) {
	t[o].l, t[o].r, t[o].todo = l, r, todoInit94
	if l == r {
		v := a[l-1]
		for i := 1; i <= v; i *= 10 {
			t[o].sum[v/i%10] += i
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg94) update(o, l, r, x, y int) {
	if l <= t[o].l && t[o].r <= r {
		v := todoInit94
		v[x] = y
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, x, y)
	}
	if m < r {
		t.update(o<<1|1, l, r, x, y)
	}
	t.maintain(o)
}

func (t seg94) maintain(o int) {
	t[o].sum = t.mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t seg94) query(o, l, r int) [10]int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf794F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, l, r, x, y int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg94, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			Fscan(in, &x, &y)
			t.update(1, l, r, x, y)
		} else {
			ans := 0
			for i, c := range t.query(1, l, r) {
				ans += i * c
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { cf794F(os.Stdin, os.Stdout) }
