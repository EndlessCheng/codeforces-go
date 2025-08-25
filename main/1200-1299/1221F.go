package main

import (
	"bufio"
	"cmp"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type pair struct{ mx, i int }
type seg []struct {
	pair
	l, r, todo int
}

func (t seg) apply(o, f int) {
	t[o].mx += f
	t[o].todo += f
}

func merge(a, b pair) pair {
	if a.mx > b.mx {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	t[o].pair = merge(t[o<<1].pair, t[o<<1|1].pair)
}

func (t seg) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mx = a[l]
		t[o].i = l
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) pair {
	if l <= t[o].l && t[o].r <= r {
		return t[o].pair
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1221F(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	// 左下角 (a,a)   右上角 (b,b)
	// a<=x<=b
	// a<=y<=b
	// a <= min(x,y) <= max(x,y) <= b
	// a <= x'
	// y' <= b
	// 按照 b 分组  sum - b + a = sum+a  - b 
	// 当我们添加 x' 时，把 [0,x'] 中的数都增加 c  搞定！
	// 最后求一下 [0,b] 中的最大值

	type point struct{ x, y, c int }
	a := make([]point, n)
	b := make([]int, 0, n*2)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y, &a[i].c)
		if a[i].x > a[i].y {
			a[i].x, a[i].y = a[i].y, a[i].x
		}
		b = append(b, a[i].x, a[i].y)
	}
	slices.SortFunc(a, func(a, b point) int { return cmp.Or(a.y-b.y, a.c-b.c) })
	slices.Sort(b)
	b = slices.Compact(b)
	m := len(b)

	x1, x2 := int(2e9), int(2e9)
	t := make(seg, 2<<bits.Len(uint(m-1)))
	t.build(b, 1, 0, m-1)
	for _, p := range a {
		t.update(1, 0, sort.SearchInts(b, p.x), p.c)
		res := t.query(1, 0, sort.SearchInts(b, p.y))
		if res.mx-p.y > ans {
			ans, x1, x2 = res.mx-p.y, b[res.i], p.y
		}
	}
	Fprintln(out, ans)
	Fprint(out, x1, x1, x2, x2)
}

func main() { cf1221F(bufio.NewReader(os.Stdin), os.Stdout) }
