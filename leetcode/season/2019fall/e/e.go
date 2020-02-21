package main

const mod int = 1e9 + 7

type lazySTNode struct {
	l, r        int
	sum         int
	addChildren int
}
type lazySegmentTree []lazySTNode

func (t lazySegmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = (lo.sum + ro.sum) % mod
}

func (t lazySegmentTree) _build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
}

func (t lazySegmentTree) _spread(o int) {
	if add := t[o].addChildren; add != 0 {
		lo, ro := &t[o<<1], &t[o<<1|1]
		lo.sum = (lo.sum + add*(lo.r-lo.l+1)) % mod
		ro.sum = (ro.sum + add*(ro.r-ro.l+1)) % mod
		lo.addChildren = (lo.addChildren + add) % mod
		ro.addChildren = (ro.addChildren + add) % mod
		t[o].addChildren = 0
	}
}

func (t lazySegmentTree) _update(o, l, r int, add int) {
	ol, or := t[o].l, t[o].r
	if l <= ol && or <= r {
		t[o].sum = (t[o].sum + add*(or-ol+1)) % mod
		t[o].addChildren = (t[o].addChildren + add) % mod
		return
	}
	t._spread(o)
	m := (ol + or) >> 1
	if l <= m {
		t._update(o<<1, l, r, add)
	}
	if m < r {
		t._update(o<<1|1, l, r, add)
	}
	t._pushUp(o)
}

func (t lazySegmentTree) _query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t._spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		res += t._query(o<<1, l, r)
	}
	if m < r {
		res += t._query(o<<1|1, l, r)
	}
	return res % mod
}

func (t lazySegmentTree) init(n int)           { t._build(1, 1, n) }
func (t lazySegmentTree) update(l, r, val int) { t._update(1, l, r, val) }
func (t lazySegmentTree) query(l, r int) int   { return t._query(1, l, r) }

func bonus(n int, leadership [][]int, operations [][]int) (ans []int) {
	g := make([][]int, n)
	for _, e := range leadership {
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
	}

	type node struct{ size, dfn int }
	nodes := make([]node, n)
	dfn := 0
	var buildNode func(int) int
	buildNode = func(v int) int {
		dfn++
		nodes[v] = node{1, dfn}
		o := &nodes[v]
		for _, w := range g[v] {
			o.size += buildNode(w)
		}
		return o.size
	}
	buildNode(0)

	t := make(lazySegmentTree, 4*n)
	t.init(n)
	for _, op := range operations {
		o := nodes[op[1]-1]
		switch op[0] {
		case 1:
			t.update(o.dfn, o.dfn, op[2])
		case 2:
			t.update(o.dfn, o.dfn+o.size-1, op[2])
		default:
			ans = append(ans, t.query(o.dfn, o.dfn+o.size-1))
		}
	}
	return
}
