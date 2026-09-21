package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick834 []int

func (t fenwick834) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick834) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return res
}

func (t fenwick834) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p3834(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	idx := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		idx[i] = i
	}
	type query struct{ l, r, k int }
	qs := make([]query, m)
	qIdx := make([]int, m)
	for i := range qs {
		qIdx[i] = i
		Fscan(in, &qs[i].l, &qs[i].r, &qs[i].k)
	}

	sorted := slices.Clone(a)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	t := make(fenwick834, n+1)

	var solve func([]int, []int, int, int)
	solve = func(idx, qIdx []int, low, high int) {
		if low == high || len(idx) <= 1 || len(qIdx) == 0 {
			for _, i := range qIdx {
				qs[i].k = a[idx[0]]
			}
			return
		}

		mid := (low + high) >> 1
		x := sorted[mid]

		var b, c []int
		for _, i := range idx {
			if a[i] <= x {
				b = append(b, i)
				t.update(i+1, 1)
			} else {
				c = append(c, i)
			}
		}

		var d, e []int
		for _, i := range qIdx {
			q := &qs[i]
			cnt := t.query(q.l, q.r)
			if cnt >= q.k {
				d = append(d, i)
			} else {
				q.k -= cnt
				e = append(e, i)
			}
		}

		for _, i := range idx {
			if a[i] <= x {
				t.update(i+1, -1)
			}
		}

		solve(b, d, low, mid)
		solve(c, e, mid+1, high)
	}

	solve(idx, qIdx, 0, len(sorted)-1)
	for _, q := range qs {
		Fprintln(out, q.k)
	}
}

//func main() { p3834(bufio.NewReader(os.Stdin), os.Stdout) }
