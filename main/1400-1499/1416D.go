package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type info16 struct{ max, i int }

func mergeInfo16(a, b info16) info16 {
	if a.max >= b.max {
		return a
	}
	return b
}

type seg16 []struct {
	l, r int
	val  info16
}

func (t seg16) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = info16{a[l], l}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg16) set0(o, i int) {
	if t[o].l == t[o].r {
		t[o].val.max = 0
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.set0(o<<1, i)
	} else {
		t.set0(o<<1|1, i)
	}
	t.maintain(o)
}

func (t seg16) maintain(o int) {
	t[o].val = mergeInfo16(t[o<<1].val, t[o<<1|1].val)
}

func (t seg16) query(o, l, r int) info16 {
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
	return mergeInfo16(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1416D(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n, m, q := r(), r(), r()
	a := make([]int, n)
	for i := range a {
		a[i] = r()
	}
	es := make([]struct{ v, w int }, m)
	for i := range es {
		es[i].v = r() - 1
		es[i].w = r() - 1
	}
	qs := make([]struct{ tp, v int }, q)
	del := make([]bool, m)
	for i := range qs {
		qs[i].tp = r()
		qs[i].v = r() - 1
		if qs[i].tp == 2 {
			del[qs[i].v] = true
		}
	}

	g := make([][]int, n*2)
	fa := make([]int, n*2)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(v, w int) {
		v = find(v)
		w = find(w)
		if v == w {
			return
		}
		fa[v] = n
		fa[w] = n
		g[n] = append(g[n], v, w)
		n++
	}
	for i, d := range del {
		if !d {
			merge(es[i].v, es[i].w)
		}
	}

	for i := q - 1; i >= 0; i-- {
		p := &qs[i]
		if p.tp == 1 {
			p.v = find(p.v)
		} else {
			e := es[p.v]
			merge(e.v, e.w)
		}
	}

	nodes := make([]struct{ in, out int }, n)
	at := make([]int, n)
	clock := -1
	var dfs func(int)
	dfs = func(v int) {
		clock++
		if v < len(a) {
			at[clock] = a[v]
		}
		nodes[v].in = clock
		for _, w := range g[v] {
			dfs(w)
		}
		nodes[v].out = clock
	}
	for i := range nodes {
		if find(i) == i { // 根
			dfs(i)
		}
	}

	t := make(seg16, 2<<bits.Len(uint(n-1)))
	t.build(at, 1, 0, n-1)
	for _, p := range qs {
		if p.tp == 2 {
			continue
		}
		node := nodes[p.v]
		res := t.query(1, node.in, node.out)
		Fprintln(out, res.max)
		if res.max > 0 {
			t.set0(1, res.i)
		}
	}
}

//func main() { cf1416D(os.Stdin, os.Stdout) }
