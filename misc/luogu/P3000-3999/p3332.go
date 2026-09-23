package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type fenwickDiff332 [][2]int

func (t fenwickDiff332) _add(i, val int) {
	for iv := i * val; i < len(t); i += i & -i {
		t[i][0] += val
		t[i][1] += iv
	}
}

// a[l] 到 a[r] 增加 val
// 1<=l<=r<=n
func (t fenwickDiff332) add(l, r, val int) {
	t._add(l, val)
	t._add(r+1, -val)
}

func (t fenwickDiff332) pre(i0 int) int {
	var s0, s1 int
	for i := i0; i > 0; i &= i - 1 {
		s0 += t[i][0]
		s1 += t[i][1]
	}
	return (i0+1)*s0 - s1
}

// 求区间和 a[l] + ... + a[r]
// 1<=l<=r<=n
func (t fenwickDiff332) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func p3332(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, l, r, k int
	Fscan(in, &n, &m)
	type query struct{ l, r, k int }
	qs := make([]query, m)
	idx := make([]int, m)
	for i := range qs {
		Fscan(in, &op, &l, &r, &k)
		if op == 1 {
			l = -l
		}
		qs[i] = query{l, r, k}
		idx[i] = i
	}

	t := make(fenwickDiff332, n+1)

	var solve func([]int, int, int)
	solve = func(idx []int, low, high int) {
		for _, i := range idx {
			if qs[i].l > 0 {
				goto next
			}
		}
		return

	next:
		if low+1 == high { // 开区间为空
			for _, i := range idx {
				if qs[i].l > 0 {
					qs[i].l = low + n
				}
			}
			return
		}

		mid := (low + high) >> 1
		var b, c []int

		for _, i := range idx {
			q := &qs[i]
			if q.l < 0 {
				if q.k < mid {
					b = append(b, i)
				} else {
					c = append(c, i)
					t.add(-q.l, q.r, 1)
				}
			} else {
				cnt := t.query(q.l, q.r)
				if cnt < q.k {
					q.k -= cnt
					b = append(b, i)
				} else {
					c = append(c, i)
				}
			}
		}

		for _, i := range idx {
			q := qs[i]
			if q.l < 0 && q.k >= mid {
				t.add(-q.l, q.r, -1)
			}
		}

		solve(b, low, mid)
		solve(c, mid, high)
	}

	// 开区间二分
	solve(idx, -n, n+1)
	for _, q := range qs {
		if q.l >= 0 {
			Fprintln(out, q.l-n)
		}
	}
}

//func main() { p3332(bufio.NewReader(os.Stdin), os.Stdout) }
