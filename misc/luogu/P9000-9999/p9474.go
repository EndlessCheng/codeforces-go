package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
type data474 struct{ l, r, pre, suf, ans int }
type seg474 []data474

func (t seg474) merge(lo, ro data474) (d data474) {
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

func (t seg474) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg474) update(o, i, v int) {
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

func p9474(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k, v, l int
	Fscan(in, &n, &k)
	t := make(seg474, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)
	type pair struct{ v, i int }
	pos := map[int][]int{}
	for i := range n {
		Fscan(in, &v)
		pos[v] = append(pos[v], i)
	}
	a := make([]int, 0, len(pos))
	for v := range pos {
		a = append(a, v)
	}
	slices.Sort(a)

	ans := int(1e9)
	for _, v := range a {
		for _, i := range pos[v] {
			t.update(1, i, 1)
		}
		for t[1].ans >= k {
			ans = min(ans, v-a[l])
			for _, i := range pos[a[l]] {
				t.update(1, i, 0)
			}
			l++
		}
	}
	Fprint(out, ans)
}

//func main() { p9474(bufio.NewReader(os.Stdin), os.Stdout) }
