package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick157 []int

func (t fenwick157) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick157) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func p3157(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v, inv int
	Fscan(in, &n, &m)
	type pair struct{ i, t, inv int }
	a := make([]pair, n+1)
	f := make(fenwick157, n+1)
	for i := range n {
		Fscan(in, &v)
		inv += f.pre(n) - f.pre(v)
		f.update(v, 1)
		a[v].i = i
	}
	for t := 1; t <= m; t++ {
		Fscan(in, &v)
		a[v].t = t
	}
	for i, p := range a {
		if p.t == 0 {
			a[i].t = m + 1
		}
	}

	f = make(fenwick157, m+2)
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		solve(l, mid)
		solve(mid, r)

		// v < w
		// p[v] > p[w]
		// 删 w：t[v] > t[w]
		// 删 v：t[w] > t[v]

		v := l
		for w := mid; w < r; w++ {
			for ; v < mid && a[v].i > a[w].i; v++ {
				f.update(a[v].t, 1)
			}
			a[w].inv += f.pre(m+1) - f.pre(a[w].t)
		}
		for v--; v >= l; v-- {
			f.update(a[v].t, -1)
		}

		w := r - 1
		for v := mid - 1; v >= l; v-- {
			for ; w >= mid && a[w].i < a[v].i; w-- {
				f.update(a[w].t, 1)
			}
			a[v].inv += f.pre(m+1) - f.pre(a[v].t)
		}
		for w++; w < r; w++ {
			f.update(a[w].t, -1)
		}

		slices.SortFunc(a[l:r], func(a, b pair) int { return b.i - a.i })
	}
	solve(1, n+1)

	slices.SortFunc(a[1:], func(a, b pair) int { return a.t - b.t })
	for _, p := range a[1:] {
		if p.t > m {
			break
		}
		Fprintln(out, inv)
		inv -= p.inv
	}
}

//func main() { p3157(bufio.NewReader(os.Stdin), os.Stdout) }
