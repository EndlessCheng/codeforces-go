package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg43 []struct{ l, r, val, todo int }

func (t seg43) do(o, v int) {
	t[o].val = v
	t[o].todo = v
}

func (t seg43) spread(o int) {
	if v := t[o].todo; v != -1 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = -1
	}
}

func (t seg43) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg43) update(o, l, r, v int) {
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
}

func (t seg43) query(o, i int) int {
	if t[o].l == t[o].r {
		return t[o].val
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		return t.query(o<<1, i)
	}
	return t.query(o<<1|1, i)
}

func cf343D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, dfn, q, op int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type node struct{ depth, size, hson, fa, top, dfn int }
	nodes := make([]node, n)
	var build func(int, int, int) int
	build = func(v, fa, d int) int {
		size, hsz, hson := 1, 0, -1
		for _, w := range g[v] {
			if w != fa {
				sz := build(w, v, d+1)
				size += sz
				if sz > hsz {
					hsz, hson = sz, w
				}
			}
		}
		nodes[v] = node{depth: d, size: size, hson: hson, fa: fa}
		return size
	}
	build(0, -1, 0)

	var markTop func(int, int)
	markTop = func(v, top int) {
		o := &nodes[v]
		o.top = top
		dfn++
		o.dfn = dfn
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

	t := make(seg43, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	doPath := func(v, w int) {
		ov, ow := nodes[v], nodes[w]
		for ; ov.top != ow.top; ov, ow = nodes[v], nodes[w] {
			topv, topw := nodes[ov.top], nodes[ow.top]
			if topv.depth < topw.depth {
				v, w = w, v
				ov, ow = ow, ov
				topv, topw = topw, topv
			}
			t.update(1, topv.dfn, ov.dfn, 0)
			v = topv.fa
		}
		if ov.depth > ow.depth {
			ov, ow = ow, ov
		}
		t.update(1, ov.dfn, ow.dfn, 0)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &v)
		v--
		if op == 1 {
			o := nodes[v]
			t.update(1, o.dfn, o.dfn+o.size-1, 1)
		} else if op == 2 {
			doPath(0, v)
		} else {
			Fprintln(out, t.query(1, nodes[v].dfn))
		}
	}
}

//func main() { cf343D(os.Stdin, os.Stdout) }
