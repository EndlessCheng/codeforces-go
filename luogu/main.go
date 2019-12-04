package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var mod int64

type lazySTNode struct {
	l, r        int
	sum         int64
	addChildren int64 // 子节点待更新
}
type lazySegmentTree []lazySTNode

func (t lazySegmentTree) _pushUp(o int) {
	t[o].sum = (t[o<<1].sum + t[o<<1|1].sum) % mod
}

func (t lazySegmentTree) _build(arr []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = arr[l]
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)
	t._pushUp(o)
}

func (t lazySegmentTree) _spread(o int) {
	if add := t[o].addChildren; add != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum = (lo.sum + add*int64(lo.r-lo.l+1)) % mod
		ro.sum = (ro.sum + add*int64(ro.r-ro.l+1)) % mod
		lo.addChildren = (lo.addChildren + add) % mod
		ro.addChildren = (ro.addChildren + add) % mod
		t[o].addChildren = 0
	}
}

func (t lazySegmentTree) _update(o, l, r int, add int64) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum = (t[o].sum + add*int64(or-ol+1)) % mod
		t[o].addChildren = (t[o].addChildren + add) % mod
		return
	}
	t._spread(o)
	mid := (ol + or) >> 1
	if l <= mid {
		t._update(o<<1, l, r, add)
	}
	if mid < r {
		t._update(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t lazySegmentTree) _query(o, l, r int) (res int64) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t._spread(o)
	mid := (t[o].l + t[o].r) >> 1
	if l <= mid {
		res = t._query(o<<1, l, r)
	}
	if mid < r {
		res += t._query(o<<1|1, l, r)
	}
	return res % mod
}

func (t lazySegmentTree) init(arr []int64)           { t._build(arr, 1, 1, len(arr)-1) }
func (t lazySegmentTree) update(l, r int, val int64) { t._update(1, l, r, val) }  // [l,r] 1<=l<=r<=n
func (t lazySegmentTree) query(l, r int) int64       { return t._query(1, l, r) } // [l,r] 1<=l<=r<=n

func solve(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n, q, root := read(), read(), read()-1
	mod = int64(read())
	vals := make([]int64, n)
	for i := range vals {
		vals[i] = int64(read()) % mod
	}
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		v, w := read()-1, read()-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	// 重儿子，父节点，深度，子树大小，所处重链顶点（深度最小），DFS 序（作为线段树中的编号，从 1 开始）
	type node struct{ hson, fa, depth, size, top, dfn int }
	nodes := make([]node, n)
	//idv := make([]int, n+1) // idv[nodes[v].dfn] == v

	var build func(v, fa, d int) *node
	build = func(v, fa, d int) *node {
		nodes[v] = node{hson: -1, fa: fa, depth: d, size: 1}
		o := &nodes[v]
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			son := build(w, v, d+1)
			o.size += son.size
			if o.hson == -1 || son.size > nodes[o.hson].size {
				o.hson = w
			}
		}
		return o
	}
	build(root, -1, 0)

	dfn := 0
	var decomposition func(v, fa, top int)
	decomposition = func(v, fa, top int) {
		o := &nodes[v]
		o.top = top
		dfn++
		o.dfn = dfn
		//idv[dfn] = v
		if o.hson != -1 {
			// 优先遍历重儿子，保证在同一条重链上的点的 DFS 序是连续的
			decomposition(o.hson, v, top)
			for _, w := range g[v] {
				if w != fa && w != o.hson {
					decomposition(w, v, w)
				}
			}
		}
	}
	decomposition(root, -1, root)

	t := make(lazySegmentTree, 4*n)
	// 点权值必须按照 DFS 序
	dfnVals := make([]int64, n+1)
	for i, v := range vals {
		dfnVals[nodes[i].dfn] = v
	}
	t.init(dfnVals)

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
			//v, w = w, v
			ov, ow = ow, ov
		}
		do(ov.dfn, ow.dfn)
	}
	updatePath := func(v, w int, add int64) {
		doPath(v, w, func(l, r int) { t.update(l, r, add) })
	}
	queryPath := func(v, w int) (sum int64) {
		doPath(v, w, func(l, r int) { sum = (sum + t.query(l, r)) % mod })
		return
	}
	updateSubtree := func(v int, add int64) {
		o := nodes[v]
		t.update(o.dfn, o.dfn+o.size-1, add)
	}
	querySubtree := func(v int) (sum int64) {
		o := nodes[v]
		return t.query(o.dfn, o.dfn+o.size-1)
	}

	for ; q > 0; q-- {
		op, v := read(), read()-1
		switch op {
		case 1:
			w, add := read()-1, int64(read())
			updatePath(v, w, add)
		case 2:
			w := read() - 1
			Fprintln(out, queryPath(v, w))
		case 3:
			add := int64(read())
			updateSubtree(v, add)
		default:
			Fprintln(out, querySubtree(v))
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
