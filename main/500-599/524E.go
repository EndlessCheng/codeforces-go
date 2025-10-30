package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type seg24 []struct{ l, r, min int }

func (t seg24) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg24) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.min = v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg24) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].min
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return min(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf524E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, k, q int
	Fscan(in, &n, &m, &k, &q)
	a := make([]struct{ x, y int }, k)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}
	qs := make([]struct{ x1, y1, x2, y2 int }, q)
	for i := range qs {
		Fscan(in, &qs[i].x1, &qs[i].y1, &qs[i].x2, &qs[i].y2)
	}

	ans := make([]bool, q)
	f := func() {
		xs := make([][]int, m+1)
		for _, p := range a {
			xs[p.y] = append(xs[p.y], p.x)
		}
		g := make([][]int, m+1)
		for i, q := range qs {
			g[q.y2] = append(g[q.y2], i)
		}

		t := make(seg24, 2<<bits.Len(uint(n-1)))
		t.build(1, 1, n)
		// 扫描线移动到 y2 时，查询 [x1,x2] 中的最小 y 是否 >= y1
		for y, xs := range xs {
			for _, x := range xs {
				t.update(1, x, y)
			}
			for _, i := range g[y] {
				q := qs[i]
				if t.query(1, q.x1, q.x2) >= q.y1 {
					ans[i] = true
				}
			}
		}
	}
	f()
	n, m = m, n
	for i := range a {
		a[i].x, a[i].y = a[i].y, a[i].x
	}
	for i := range qs {
		q := &qs[i]
		q.x1, q.y1 = q.y1, q.x1
		q.x2, q.y2 = q.y2, q.x2
	}
	f()

	for _, b := range ans {
		if b {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf524E(bufio.NewReader(os.Stdin), os.Stdout) }
