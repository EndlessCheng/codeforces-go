package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF383C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w, dfn, op, val int
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

	nodes := make([]struct{ sgn, l, r int }, n)
	var f func(int, int, int) int
	f = func(v, fa, sgn int) int {
		nodes[v].sgn = sgn
		dfn++
		nodes[v].l = dfn
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v, -sgn)
			}
		}
		nodes[v].r = nodes[v].l + sz
		return sz
	}
	f(0, -1, 1)

	tree := make([]int, n+1)
	add := func(i, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	pre := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}

	for ; q > 0; q-- {
		Fscan(in, &op, &v)
		v--
		o := nodes[v]
		if op == 1 {
			Fscan(in, &val)
			add(o.l, val*o.sgn)
			add(o.r, -val*o.sgn)
		} else {
			Fprintln(out, a[v]+pre(o.l)*o.sgn)
		}
	}
}

//func main() { CF383C(os.Stdin, os.Stdout) }
