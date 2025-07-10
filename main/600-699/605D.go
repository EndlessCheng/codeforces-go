package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
type pair5 struct{ b, i int }
var g5 [][]pair5

type seg5 []struct{ l, r, min int }

func (t seg5) maintain(o int) {
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg5) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].min = g5[l][0].b
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg5) delete(o, qr, maxY int, f func(int)) {
	l := t[o].l
	if l >= qr || t[o].min > maxY {
		return
	}
	if l == t[o].r {
		f(l)
		t[o].min = g5[l][0].b
		return
	}
	t.delete(o<<1|1, qr, maxY, f)
	t.delete(o<<1, qr, maxY, f)
	t.maintain(o)
}

func cf605D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]struct{ a, b, c, d int }, n+1)
	xs := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].a, &a[i].b, &a[i].c, &a[i].d)
		xs[i] = a[i].a
	}
	slices.Sort(xs)
	xs = slices.Compact(xs)
	m := len(xs)

	g5 = make([][]pair5, m)
	for i := 1; i <= n; i++ {
		x := sort.SearchInts(xs, a[i].a)
		g5[x] = append(g5[x], pair5{a[i].b, i})
	}
	for i, ps := range g5 {
		slices.SortFunc(ps, func(a, b pair5) int { return a.b - b.b })
		g5[i] = append(ps, pair5{2e9, 0}) // 哨兵
	}
	t := make(seg5, 2<<bits.Len(uint(m-1)))
	t.build(1, 0, m-1)

	q := []int{0}
	pre := make([]int, n+1)
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if i == n {
			ans := []int{}
			for ; i > 0; i = pre[i] {
				ans = append(ans, i)
			}
			Fprintln(out, len(ans))
			for i := len(ans) - 1; i >= 0; i-- {
				Fprint(out, ans[i], " ")
			}
			Fprintln(out)
			return
		}

		maxY := a[i].d
		t.delete(1, sort.SearchInts(xs, a[i].c+1), maxY, func(l int) {
			ps := g5[l]
			for ps[0].b <= maxY {
				pre[ps[0].i] = i
				q = append(q, ps[0].i)
				ps = ps[1:]
			}
			g5[l] = ps
		})
	}
	Fprint(out, -1)
}

//func main() { cf605D(bufio.NewReader(os.Stdin), os.Stdout) }
