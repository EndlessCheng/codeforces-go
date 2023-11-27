package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1900E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		g := make([][]int, n)
		rg := make([][]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			rg[w] = append(rg[w], v)
		}

		vs := make([]int, 0, n)
		vis := make([]bool, n)
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
			vs = append(vs, v)
		}
		for i, b := range vis {
			if !b {
				dfs(i)
			}
		}

		scc := [][]int{}
		clear(vis)
		var comp []int
		var rdfs func(int)
		rdfs = func(v int) {
			vis[v] = true
			comp = append(comp, v)
			for _, w := range rg[v] {
				if !vis[w] {
					rdfs(w)
				}
			}
		}
		for i := n - 1; i >= 0; i-- {
			v := vs[i]
			if vis[v] {
				continue
			}
			comp = []int{}
			rdfs(v)
			scc = append(scc, comp)
		}

		ns := len(scc)
		a2 := make([]int, ns)
		sid := make([]int, n)
		for i, cc := range scc {
			for _, v := range cc {
				a2[i] += a[v]
				sid[v] = i
			}
		}

		g2 := make([][]int, ns)
		deg := make([]int, ns)
		for v, ws := range g {
			v = sid[v]
			for _, w := range ws {
				w = sid[w]
				if v != w {
					g2[v] = append(g2[v], w)
					deg[w]++
				}
			}
		}

		type pair struct{ len, s int }
		f := make([]pair, ns)
		q := make([]int, 0, ns)
		for i, d := range deg {
			if d == 0 {
				q = append(q, i)
			}
		}
		ans := pair{}
		gr := func(p, q pair) bool { return p.len > q.len || p.len == q.len && p.s < q.s }
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			p := f[v]
			p.len += len(scc[v])
			p.s += a2[v]
			if gr(p, ans) {
				ans = p
			}
			for _, w := range g2[v] {
				if gr(p, f[w]) {
					f[w] = p
				}
				if deg[w]--; deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		Fprintln(out, ans.len, ans.s)
	}
}

//func main() { CF1900E(os.Stdin, os.Stdout) }
