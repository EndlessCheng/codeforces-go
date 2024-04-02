package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg20 []struct {
	l, r int
	mask uint
	todo uint
}

func (t seg20) do(o int, v uint) {
	t[o].mask = v
	t[o].todo = v
}

func (t seg20) spread(o int) {
	if v := t[o].todo; v > 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg20) build(a []uint, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].mask = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg20) update(o, l, r int, v uint) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg20) maintain(o int) {
	t[o].mask = t[o<<1].mask | t[o<<1|1].mask
}

func (t seg20) query(o, l, r int) uint {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mask
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) | t.query(o<<1|1, l, r)
}

func cf620E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, dfn, op, v, w int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	b := make([]uint, n)
	nodes := make([]struct{ l, r int }, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) (size int) {
		b[dfn] = 1 << a[v]
		dfn++
		nodes[v].l = dfn
		for _, w := range g[v] {
			if w != fa {
				sz := dfs(w, v)
				size += sz
			}
		}
		nodes[v].r = nodes[v].l + size
		size++
		return
	}
	dfs(0, -1)

	t := make(seg20, 2<<bits.Len(uint(n-1)))
	t.build(b, 1, 1, n)
	for ; m > 0; m-- {
		Fscan(in, &op, &v)
		o := nodes[v-1]
		if op == 1 {
			Fscan(in, &w)
			t.update(1, o.l, o.r, 1<<w)
		} else {
			Fprintln(out, bits.OnesCount(t.query(1, o.l, o.r)))
		}
	}
}

//func main() { cf620E(os.Stdin, os.Stdout) }
