package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick810 []int

func (t fenwick810) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick810) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return res
}

func p3810(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	type tuple struct{ x, y, z int }
	a := make([]tuple, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].z)
	}
	slices.SortFunc(a, func(a, b tuple) int { return cmp.Or(a.x-b.x, a.y-b.y, a.z-b.z) })

	type data struct{ y, z, dup, res int }
	b := []data{}
	st := -1
	for i, t := range a {
		if i == n-1 || t != a[i+1] {
			b = append(b, data{t.y, t.z, i - st, i - st - 1})
			st = i
		}
	}

	t := make(fenwick810, k+1)
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		m := (l + r) >> 1
		solve(l, m)
		solve(m, r)

		i := l
		for j := m; j < r; j++ {
			for ; i < m && b[i].y <= b[j].y; i++ {
				t.update(b[i].z, b[i].dup)
			}
			b[j].res += t.pre(b[j].z)
		}
		for i--; i >= l; i-- {
			t.update(b[i].z, -b[i].dup)
		}

		slices.SortFunc(b[l:r], func(a, b data) int { return cmp.Or(a.y-b.y, a.z-b.z) })
	}
	solve(0, len(b))

	ans := make([]int, n)
	for _, d := range b {
		ans[d.res] += d.dup
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p3810(bufio.NewReader(os.Stdin), os.Stdout) }
