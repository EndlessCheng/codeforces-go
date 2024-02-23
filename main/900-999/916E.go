package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type fenwickDiff16 [][2]int

func (t fenwickDiff16) _add(i, val int) {
	for iv := i * val; i < len(t); i += i & -i {
		t[i][0] += val
		t[i][1] += iv
	}
}

func (t fenwickDiff16) add(l, r, val int) {
	t._add(l, val)
	t._add(r+1, -val)
}

func (t fenwickDiff16) pre(i0 int) int {
	var s0, s1 int
	for i := i0; i > 0; i &= i - 1 {
		s0 += t[i][0]
		s1 += t[i][1]
	}
	return (i0+1)*s0 - s1
}

func (t fenwickDiff16) query(l, r int) int {
	return t.pre(r) - t.pre(l-1)
}

func cf916E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, dfn, op, v, w, val, rt int
	Fscan(in, &n, &q)
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

	nodes := make([]struct{ l, r int }, n)
	const mx = 17
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var build func(int, int) int
	build = func(v, p int) int {
		dfn++
		nodes[v].l = dfn
		pa[v][0] = p
		sz := 1
		for _, w := range g[v] {
			if w != p {
				dep[w] = dep[v] + 1
				sz += build(w, v)
				a[v] += a[w]
			}
		}
		nodes[v].r = nodes[v].l + sz - 1
		return sz
	}
	build(0, -1)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	uptoDep := func(v, d int) int {
		for k := uint(dep[v] - d); k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
		}
		return v
	}
	getLCA := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		w = uptoDep(w, dep[v])
		if w == v {
			return v
		}
		for i := mx - 1; i >= 0; i-- {
			if pv, pw := pa[v][i], pa[w][i]; pv != pw {
				v, w = pv, pw
			}
		}
		return pa[v][0]
	}
	isAncestor := func(f, v int) bool { return nodes[f].l < nodes[v].l && nodes[v].l <= nodes[f].r }

	t := make(fenwickDiff16, n+1)
	for ; q > 0; q-- {
		Fscan(in, &op, &v)
		v--
		if op == 1 {
			rt = v
		} else if op == 2 {
			Fscan(in, &w, &val)
			w--
			lca := getLCA(v, w)
			if lca == rt {
				t.add(1, n, val)
			} else if !isAncestor(lca, rt) {
				// 更新 lca 子树
				p := nodes[lca]
				t.add(p.l, p.r, val)
			} else { // lca 是 rt 的祖先
				maxD := max(dep[getLCA(rt, v)], dep[getLCA(rt, w)])
				if maxD < dep[rt] {
					subV := uptoDep(rt, maxD+1)
					p := nodes[subV]
					t.add(p.l, p.r, -val)
				}
				// 更新整棵树
				t.add(1, n, val)
			}
		} else {
			if v == rt {
				Fprintln(out, a[0]+t.query(1, n))
			} else if !isAncestor(v, rt) {
				p := nodes[v]
				Fprintln(out, a[v]+t.query(p.l, p.r))
			} else { // v 是 rt 的祖先
				subV := uptoDep(rt, dep[v]+1)
				p := nodes[subV]
				Fprintln(out, a[0]+t.query(1, n)-a[subV]-t.query(p.l, p.r))
			}
		}
	}
}

//func main() { cf916E(os.Stdin, os.Stdout) }
