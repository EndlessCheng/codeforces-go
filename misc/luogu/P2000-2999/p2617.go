package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick []int

func (t fenwick) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return res
}

func (t fenwick) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p2617(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v, l, r, k int
	var op string
	Fscan(in, &n, &m)
	a := make([]int, n+1)
	type query struct{ l, r, k int }
	qs := make([]query, n, n+m*2)
	idx := make([]int, n, n+m*2)
	sorted := make([]int, n, n+m)
	for i := range n {
		Fscan(in, &v)
		a[i+1] = v
		qs[i] = query{i + 1, v, -1}
		idx[i] = i
		sorted[i] = v
	}
	for range m {
		Fscan(in, &op, &l, &r)
		if op[0] == 'Q' {
			Fscan(in, &k)
			idx = append(idx, len(qs))
			qs = append(qs, query{l, r, k})
		} else {
			idx = append(idx, len(qs), len(qs)+1)
			qs = append(qs, query{l, a[l], -3}, query{l, r, -1})
			a[l] = r
			sorted = append(sorted, r)
		}
	}

	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	t := make(fenwick, n+1)

	memB := make([]int, 0, len(idx))
	memC := make([]int, 0, len(idx))

	var solve func([]int, int, int)
	solve = func(idx []int, low, high int) {
		for _, i := range idx {
			if qs[i].k > 0 {
				goto next
			}
		}
		return

	next:
		if low == high {
			for _, i := range idx {
				if qs[i].k > 0 {
					qs[i].k = sorted[low]
				}
			}
			return
		}

		mid := (low + high) >> 1
		x := sorted[mid]

		b := memB[:0]
		c := memC[:0]
		for _, p := range idx {
			q := &qs[p]
			if q.k < 0 {
				if q.r <= x {
					b = append(b, p)
					t.update(q.l, q.k+2)
				} else {
					c = append(c, p)
				}
			} else {
				cnt := t.query(q.l, q.r)
				if cnt >= q.k {
					b = append(b, p)
				} else {
					q.k -= cnt
					c = append(c, p)
				}
			}
		}

		for _, i := range idx {
			q := qs[i]
			if q.k < 0 && q.r <= x {
				t.update(q.l, -q.k-2)
			}
		}

		nb := len(b)
		copy(idx, b)
		copy(idx[nb:], c)
		solve(idx[:nb], low, mid)
		solve(idx[nb:], mid+1, high)
	}

	solve(idx, 0, len(sorted)-1)
	for _, q := range qs[n:] {
		if q.k >= 0 {
			Fprintln(out, q.k)
		}
	}
}

//func main() { p2617(bufio.NewReader(os.Stdin), os.Stdout) }
