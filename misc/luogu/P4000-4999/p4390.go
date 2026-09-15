package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick390 []int

func (t fenwick390) update(i, val int) {
	for ; i < len(t); i += i & -i {
		t[i] += val
	}
}

func (t fenwick390) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

func p4390(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var mx, op, q, x1, y1, x2, y2 int
	Fscan(in, &mx, &mx)
	type data struct{ x, y, c, res int }
	a := make([]data, 0, 2e5)
	for {
		Fscan(in, &op, &x1, &y1, &x2)
		x1++
		if op == 3 {
			break
		}
		if op == 1 {
			a = append(a, data{x1, y1, x2, -1})
		} else {
			x2++
			Fscan(in, &y2)
			q++
			a = append(a,
				data{x1 - 1, y1 - 1, q, 0},
				data{x1 - 1, y2, -q, 0},
				data{x2, y1 - 1, -q, 0},
				data{x2, y2, q, 0})
		}
	}

	t := make(fenwick390, mx+2)
	var solve func(int, int)
	solve = func(l, r int) {
		if l+1 == r {
			return
		}
		mid := (l + r) >> 1
		solve(l, mid)
		solve(mid, r)

		i := l
		for j := mid; j < r; j++ {
			for ; i < mid && a[i].y <= a[j].y; i++ {
				if a[i].res < 0 {
					t.update(a[i].x, a[i].c)
				}
			}
			if a[j].res >= 0 {
				a[j].res += t.pre(a[j].x)
			}
		}
		for i--; i >= l; i-- {
			if a[i].res < 0 {
				t.update(a[i].x, -a[i].c)
			}
		}

		slices.SortFunc(a[l:r], func(a, b data) int { return a.y - b.y })
	}
	solve(0, len(a))

	ans := make([]int, q)
	for _, d := range a {
		if d.res > 0 {
			if d.c > 0 {
				ans[d.c-1] += d.res
			} else {
				ans[-d.c-1] -= d.res
			}
		}
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() { p4390(bufio.NewReader(os.Stdin), os.Stdout) }
