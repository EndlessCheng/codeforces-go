package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg8 []struct{ l, r, min, todo int }

func (t seg8) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o].min = min8(t[o<<1].min, t[o<<1|1].min)
}

func (t seg8) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg8) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg8) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t[o].min = min8(t[o<<1].min, t[o<<1|1].min)
}

func CF1108E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	t := make(seg8, n*4)
	t.build(a, 1, 1, n)

	ps := make([]struct{ l, r int }, m)
	ls := make([][]int, n+1)
	for i := range ps {
		Fscan(in, &ps[i].l, &ps[i].r)
		l, r := ps[i].l, ps[i].r
		t.update(1, l, r, -1)
		ls[l] = append(ls[l], r)
	}

	maxD, maxI := 0, 1
	rs := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		for _, r := range ls[i] {
			t.update(1, i, r, 1)
			rs[r] = append(rs[r], i)
		}
		for _, l := range rs[i-1] {
			t.update(1, l, i-1, -1)
		}
		d := a[i] - t[1].min
		if d > maxD {
			maxD, maxI = d, i
		}
	}

	Fprintln(out, maxD)
	ids := []interface{}{}
	for i, p := range ps {
		if p.r < maxI || p.l > maxI {
			ids = append(ids, i+1)
		}
	}
	Fprintln(out, len(ids))
	Fprintln(out, ids...)
}

//func main() { CF1108E2(os.Stdin, os.Stdout) }

func min8(a, b int) int {
	if a > b {
		return b
	}
	return a
}
