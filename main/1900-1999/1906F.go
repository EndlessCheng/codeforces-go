package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type info6 struct{ ans, tot, pre, suf int }

type seg6 []struct {
	l, r int
	info6
}

func (t seg6) mergeInfo(a, b info6) info6 {
	return info6{
		max(a.ans, b.ans, a.suf+b.pre),
		a.tot + b.tot,
		max(a.pre, a.tot+b.pre),
		max(b.suf, b.tot+a.suf),
	}
}

func (t seg6) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg6) update(o, i, v int) {
	if t[o].l == t[o].r {
		v += t[o].info6.tot
		t[o].info6 = info6{v, v, v, v}
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o].info6 = t.mergeInfo(t[o<<1].info6, t[o<<1|1].info6)
}

func (t seg6) query(o, l, r int) (d info6) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].info6
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1906F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, l, r, v int
	Fscan(in, &n, &m)
	type pair struct{ i, v int }
	ops := make([][]pair, n+2)
	for i := 1; i <= m; i++ {
		Fscan(in, &l, &r, &v)
		ops[l] = append(ops[l], pair{i, v})
		ops[r+1] = append(ops[r+1], pair{i, -v})
	}

	Fscan(in, &q)
	type query struct{ l, r, qi int }
	qs := make([][]query, n+1)
	for i := range q {
		Fscan(in, &v, &l, &r)
		qs[v] = append(qs[v], query{l, r, i})
	}

	ans := make([]int, q)
	t := make(seg6, 2<<bits.Len(uint(m)))
	t.build(1, 1, m)
	for i, qs := range qs {
		for _, p := range ops[i] {
			t.update(1, p.i, p.v)
		}
		for _, q := range qs {
			ans[q.qi] = t.query(1, q.l, q.r).ans
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { cf1906F(bufio.NewReader(os.Stdin), os.Stdout) }
