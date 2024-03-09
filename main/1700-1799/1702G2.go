package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1702G2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w, dfn, Q, k, aq int
	Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	nodes := make([]struct{ l, r int }, n)
	const mx = 18
	pa := make([][mx]int, n)
	dep := make([]uint, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) (size int) {
		pa[v][0] = fa
		dfn++
		nodes[v].l = dfn
		for _, w := range g[v] {
			if w != fa {
				dep[w] = dep[v] + 1
				sz := dfs(w, v)
				size += sz
			}
		}
		size++
		nodes[v].r = nodes[v].l + size
		return
	}
	dfs(0, -1)
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p != -1 {
				pa[v][i+1] = pa[p][i]
			} else {
				pa[v][i+1] = -1
			}
		}
	}
	isAncestor := func(f, v int) bool { return nodes[f].l < nodes[v].l && nodes[v].l < nodes[f].r }
	up := func(v int, d uint) int {
		for k := dep[v] - d; k > 0; k &= k - 1 {
			v = pa[v][bits.TrailingZeros(k)]
		}
		return v
	}

o:
	for Fscan(in, &Q); Q > 0; Q-- {
		Fscan(in, &k)
		a := make([]int, k)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
			if dep[a[i]] < dep[a[0]] {
				a[i], a[0] = a[0], a[i]
			}
		}
		if k <= 2 {
			Fprintln(out, "YES")
			continue
		}
		p, q := a[0], a[1] // p 的深度最小
		top := isAncestor(p, q)
		if top {
			aq = up(q, dep[p]+1) // aq 是 p 的儿子和 q 的祖先
		}
		for _, v := range a[2:] {
			if top {
				if isAncestor(q, v) { // v 在 q 下面
					q = v
				} else if !isAncestor(v, q) { // v 不在 p 和 q 之间
					if isAncestor(aq, v) { // v 在 p 和 q 之间的分叉上
						Fprintln(out, "NO")
						continue o
					}
					p = v
					top = false
				}
			} else if isAncestor(p, v) { // v 在 p 下面
				p = v
			} else if isAncestor(q, v) { // v 在 q 下面
				q = v
			} else if !isAncestor(v, p) && !isAncestor(v, q) { // v 不在 p 到 q 的路径上
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1702G2(os.Stdin, os.Stdout) }
