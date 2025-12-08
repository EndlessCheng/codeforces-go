package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg137 []struct{ l, r, mn int }

func (t seg137) maintain(o int) {
	t[o].mn = min(t[o<<1].mn, t[o<<1|1].mn)
}

func (t seg137) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].mn = -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg137) update(o, i, val int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.mn = val
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg137) query(o, ql int) int {
	l, r := t[o].l, t[o].r
	if l == r {
		if t[o].mn >= ql {
			return l + 1
		}
		return l
	}
	if t[o<<1].mn < ql {
		return t.query(o<<1, ql)
	}
	return t.query(o<<1|1, ql)
}

func p4137Offline(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type pair struct{ l, i int }
	g := make([][]pair, n)
	for i := range m {
		Fscan(in, &l, &r)
		l--
		r--
		g[r] = append(g[r], pair{l, i})
	}

	ans := make([]int, m)
	t := make(seg137, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	for i, v := range a {
		if v < n {
			t.update(1, v, i)
		}
		for _, p := range g[i] {
			ans[p.i] = t.query(1, p.l)
		}
	}

	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p4137Offline(bufio.NewReader(os.Stdin), os.Stdout) }
