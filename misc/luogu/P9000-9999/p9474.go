package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
type data struct{ l, r, pre, suf, ans int }
type seg []data

func (t seg) merge(lo, ro data) (d data) {
	d.l = lo.l
	d.r = ro.r

	d.pre = lo.pre
	if lo.pre == lo.r-lo.l+1 {
		d.pre += ro.pre
	}

	d.suf = ro.suf
	if ro.suf == ro.r-ro.l+1 {
		d.suf += lo.suf
	}

	d.ans = lo.ans + ro.ans - (lo.suf+1)/2 - (ro.pre+1)/2 + (lo.suf+ro.pre+1)/2
	return
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.pre = v
		cur.suf = v
		cur.ans = v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t[o] = t.merge(t[o<<1], t[o<<1|1])
}

func p9474(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	slices.SortFunc(a, func(a, b pair) int { return a.v - b.v })

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)

	ans := int(1e9)
	l := 0
	for _, p := range a {
		t.update(1, p.i, 1)
		for t[1].ans >= m {
			ans = min(ans, p.v-a[l].v)
			t.update(1, a[l].i, 0)
			l++
		}
	}
	Fprint(out, ans)
}

//func main() { p9474(bufio.NewReader(os.Stdin), os.Stdout) }
