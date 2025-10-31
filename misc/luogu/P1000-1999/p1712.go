package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
type seg712 []struct{ l, r, mx, todo int }

func (t seg712) apply(o, f int) {
	cur := &t[o]
	cur.mx += f
	cur.todo += f
}

func (t seg712) maintain(o int) {
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg712) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg712) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg712) update(o, l, r, f int) {
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

func p1712(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	type pair struct{ l, r int }
	a := make([]pair, n)
	b := make([]int, 0, n*2)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r)
		b = append(b, a[i].l, a[i].r)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.r - a.l - (b.r - b.l) })
	slices.Sort(b)
	b = slices.Compact(b)
	k := len(b)

	t := make(seg712, 2<<bits.Len(uint(k-1)))
	t.build(1, 0, k-1)
	ans := int(2e9)
	left := 0
	for _, p := range a {
		t.update(1, sort.SearchInts(b, p.l), sort.SearchInts(b, p.r), 1)
		for t[1].mx >= m {
			q := a[left]
			ans = min(ans, p.r-p.l-(q.r-q.l))
			t.update(1, sort.SearchInts(b, q.l), sort.SearchInts(b, q.r), -1)
			left++
		}
	}
	if ans == 2e9 {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { p1712(bufio.NewReader(os.Stdin), os.Stdout) }
