package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type seg22 []struct{ l, r, val int }

func (t seg22) build(o, l, r int) {
	t[o].l, t[o].r, t[o].val = l, r, 1e9
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg22) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].val = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t[o].val = min22(t[o<<1].val, t[o<<1|1].val)
}

func (t seg22) query(o, l, r int) int {
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
	return min22(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF522D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, l, r int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	type pair struct{ l, i int }
	q := make([][]pair, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &l, &r)
		q[r] = append(q[r], pair{l, i})
	}

	ans := make([]int, m)
	t := make(seg22, n*4)
	t.build(1, 1, n)
	last := map[int]int{}
	for i := 1; i <= n; i++ {
		if p := last[a[i]]; p > 0 {
			t.update(1, p, i-p)
		}
		for _, p := range q[i] {
			res := t.query(1, p.l, i)
			if res == 1e9 {
				res = -1
			}
			ans[p.i] = res
		}
		last[a[i]] = i
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { CF522D(os.Stdin, os.Stdout) }

func min22(a, b int) int {
	if a > b {
		return b
	}
	return a
}
