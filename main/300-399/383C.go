package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF383C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w, op, val int
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

	tree := make([]int, n+1)
	add := func(i, v int) {
		for ; i <= n; i += i & -i {
			tree[i] += v
		}
	}
	addR := func(l, r, v int) { add(l, v); add(r, -v) }
	query := func(i int) (v int) {
		for ; i > 0; i &= i - 1 {
			v += tree[i]
		}
		return
	}

	type node struct{ dfn, sz, d int }
	nodes := make([]node, n)
	dfn := 0
	var f func(v, fa, d int) int
	f = func(v, fa, d int) int {
		dfn++
		nodes[v].dfn = dfn
		nodes[v].d = d
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v, d+1)
			}
		}
		nodes[v].sz = sz
		return sz
	}
	f(0, -1, 0)

	for ; q > 0; q-- {
		Fscan(in, &op, &v)
		v--
		o := nodes[v]
		if op == 1 {
			Fscan(in, &val)
			if o.d&1 > 0 {
				val = -val
			}
			addR(o.dfn, o.dfn+o.sz, val)
		} else {
			ans := query(o.dfn)
			if o.d&1 > 0 {
				ans = -ans
			}
			Fprintln(out, a[v]+ans)
		}
	}
}

//func main() { CF383C(os.Stdin, os.Stdout) }
