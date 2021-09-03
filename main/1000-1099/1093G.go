package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg93 []struct {
	l, r int
	a    []int
}

func (t seg93) set(o int, a []int) {
	for i := range t[o].a {
		t[o].a[i] = 0
		for j, v := range a {
			t[o].a[i] += v * (i>>j&1<<1 - 1)
		}
	}
}

func (seg93) op(a, b []int) []int {
	c := make([]int, len(a))
	for i, v := range a {
		c[i] = max93(v, b[i])
	}
	return c
}

func (t seg93) maintain(o int) {
	t[o].a = t.op(t[o<<1].a, t[o<<1|1].a)
}

func (t seg93) build(a [][]int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].a = make([]int, 1<<len(a[l-1]))
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg93) update(o, i int, a []int) {
	if t[o].l == t[o].r {
		t.set(o, a)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, a)
	} else {
		t.update(o<<1|1, i, a)
	}
	t.maintain(o)
}

func (t seg93) query(o, l, r int) []int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].a
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.op(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF1093G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, q, op, l, r int
	Fscan(in, &n, &k)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, k)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}
	t := make(seg93, n*4)
	t.build(a, 1, 1, n)
	b := make([]int, k)
	m := 1<<k - 1
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op, &l); op == 1 {
			for i := range b {
				Fscan(in, &b[i])
			}
			t.update(1, l, b)
		} else {
			Fscan(in, &r)
			ans := -int(1e9)
			res := t.query(1, l, r)
			for i, v := range res {
				ans = max93(ans, v+res[m^i])
			}
			Fprintln(out, ans)
		}
	}
}

func max93(a, b int) int {
	if b > a {
		return b
	}
	return a
}

//func main() { CF1093G(os.Stdin, os.Stdout) }
