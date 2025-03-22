package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg3178 []struct {
	l, r, sum, todo int
}

func (t seg3178) apply(o, f int) {
	cur := &t[o]
	cur.sum += f * (cur.r - cur.l + 1)
	cur.todo += f
}

func (t seg3178) maintain(o int) {
	t[o].sum = t[o<<1].sum + t[o<<1|1].sum
}

func (t seg3178) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg3178) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg3178) update(o, l, r, f int) {
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

func (t seg3178) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func p3178(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, x, add int
	Fscan(in, &n, &m)
	vals := make([]int, n)
	for i := range vals {
		Fscan(in, &vals[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type node struct{ depth, size, hson, fa, top, dfn int }
	nodes := make([]node, n)
	var build func(int, int, int) int
	build = func(v, fa, dep int) int {
		size, hsz, hson := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				sz := build(w, v, dep+1)
				size += sz
				if sz > hsz {
					hsz, hson = sz, w
				}
			}
		}
		nodes[v] = node{depth: dep, size: size, hson: hson, fa: fa}
		return size
	}
	build(0, -1, 0)

	dfn := 0
	var markTop func(int, int)
	markTop = func(v, top int) {
		o := &nodes[v]
		o.top = top
		o.dfn = dfn
		dfn++
		if o.hson != -1 {
			markTop(o.hson, top)
			for _, w := range g[v] {
				if w != o.fa && w != o.hson {
					markTop(w, w)
				}
			}
		}
	}
	markTop(0, 0)

	dfnVals := make([]int, n)
	for i, v := range vals {
		dfnVals[nodes[i].dfn] = v
	}
	t := make(seg3178, 2<<bits.Len(uint(n-1)))
	t.build(dfnVals, 1, 0, n-1)

	doPath := func(v, w int, do func(l, r int)) {
		ov, ow := nodes[v], nodes[w]
		for ; ov.top != ow.top; ov, ow = nodes[v], nodes[w] {
			topv, topw := nodes[ov.top], nodes[ow.top]
			if topv.depth < topw.depth {
				v, w = w, v
				ov, ow = ow, ov
				topv, topw = topw, topv
			}
			do(topv.dfn, ov.dfn)
			v = topv.fa
		}
		if ov.depth > ow.depth {
			ov, ow = ow, ov
		}
		do(ov.dfn, ow.dfn)
	}
	updatePath := func(v, w, add int) { doPath(v, w, func(l, r int) { t.update(1, l, r, add) }) }
	updateSubtree := func(v, add int) { o := nodes[v]; t.update(1, o.dfn, o.dfn+o.size-1, add) }
	queryPath := func(v, w int) (s int) { doPath(v, w, func(l, r int) { s += t.query(1, l, r) }); return }

	for ; m > 0; m-- {
		Fscan(in, &op, &x)
		x--
		if op == 1 {
			Fscan(in, &add)
			updatePath(x, x, add)
		} else if op == 2 {
			Fscan(in, &add)
			updateSubtree(x, add)
		} else {
			Fprintln(out, queryPath(0, x))
		}
	}
}

//func main() { p3178(bufio.NewReader(os.Stdin), os.Stdout) }
