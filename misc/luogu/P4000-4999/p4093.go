package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick093 []int

func (t fenwick093) reset(i int) {
	for ; i < len(t); i += i & -i {
		t[i] = 0
	}
}

func (t fenwick093) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] = max(t[i], val)
	}
}

func (t fenwick093) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, t[i])
	}
	return
}

func p4093(in io.Reader, out io.Writer) {
	var n, m, p, v int
	Fscan(in, &n, &m)
	type data struct{ i, v, mn, mx, f int }
	d := make([]data, n)
	for i := range d {
		Fscan(in, &v)
		d[i] = data{i, v, v, v, 1}
	}
	for range m {
		Fscan(in, &p, &v)
		p--
		d[p].mn = min(d[p].mn, v)
		d[p].mx = max(d[p].mx, v)
	}

	// f[i] = max(f[j]) + 1
	// j < i
	// mx[j] <= a[i]
	// a[j] <= mn[i]

	t := make(fenwick093, 1e5+1)
	ans := 1
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		solve(l, mid)

		slices.SortFunc(d[l:mid], func(a, b data) int { return a.mx - b.mx })
		slices.SortFunc(d[mid:r], func(a, b data) int { return a.v - b.v })
		j := l
		for i := mid; i < r; i++ {
			for ; j < mid && d[j].mx <= d[i].v; j++ {
				t.update(d[j].v, d[j].f)
			}
			d[i].f = max(d[i].f, t.pre(d[i].mn)+1)
			ans = max(ans, d[i].f)
		}
		for j--; j >= l; j-- {
			t.reset(d[j].v)
		}
		slices.SortFunc(d[l:r], func(a, b data) int { return a.i - b.i })

		solve(mid, r)
	}
	solve(0, n)
	Fprint(out, ans)
}

//func main() { p4093(bufio.NewReader(os.Stdin), os.Stdout) }
