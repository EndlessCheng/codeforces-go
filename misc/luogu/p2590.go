package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type data2590 struct{ mx, s int }

type seg2590 []struct {
	l, r int
	val  data2590
}

func mergeInfo2590(a, b data2590) data2590 {
	return data2590{max(a.mx, b.mx), a.s + b.s}
}

func (t seg2590) maintain(o int) {
	t[o].val = mergeInfo2590(t[o<<1].val, t[o<<1|1].val)
}

func (t seg2590) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = data2590{a[l], a[l]}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg2590) update(o, i, v int) {
	if t[o].l == t[o].r {
		t[o].val = data2590{v, v}
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t seg2590) query(o, l, r int) data2590 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo2590(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p2590(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q int
	var op string
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	// 深度，子树大小，重儿子，父节点，所处重链顶点（深度最小），DFS 序（作为线段树中的编号）
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
			// 优先遍历重儿子，保证在同一条重链上的点的 DFS 序是连续的
			markTop(o.hson, top)
			for _, w := range g[v] {
				if w != o.fa && w != o.hson {
					markTop(w, w)
				}
			}
		}
	}
	markTop(0, 0)

	// 按照 DFS 序对应的点权初始化线段树
	dfnVals := make([]int, n)
	for i, v := range a {
		dfnVals[nodes[i].dfn] = v
	}
	t := make(seg2590, 2<<bits.Len(uint(n-1)))
	t.build(dfnVals, 1, 0, n-1)

	doPath := func(v, w int, do func(l, r int)) {
		ov, ow := nodes[v], nodes[w]
		for ; ov.top != ow.top; ov, ow = nodes[v], nodes[w] {
			topv, topw := nodes[ov.top], nodes[ow.top]
			// v 所处的重链顶点必须比 w 的深
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
	updatePath := func(v, w, val int) { doPath(v, w, func(l, r int) { t.update(1, l, val) }) }
	queryPath := func(v, w int) (res data2590) { res.mx = -1e18; doPath(v, w, func(l, r int) { res = mergeInfo2590(res, t.query(1, l, r)) }); return }

	for Fscan(in, &q); q > 0; q-- {
		var v, w int
		Fscan(in, &op, &v, &w)
		v--
		w--
		if op[1] == 'H' {
			updatePath(v, v, w+1)
		} else if op[1] == 'M' {
			Fprintln(out, queryPath(v, w).mx)
		} else {
			Fprintln(out, queryPath(v, w).s)
		}
	}
}

//func main() { p2590(bufio.NewReader(os.Stdin), os.Stdout) }
