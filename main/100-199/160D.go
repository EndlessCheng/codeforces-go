package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF160D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, v, w, wi, c int
	Fscan(in, &n, &m)
	fa := make([]int, n+1)
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
	type edge struct{ v, w, i int }
	edges := [1e6 + 1][]edge{}
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w, &wi)
		edges[wi] = append(edges[wi], edge{v, w, i})
	}
	ans := make([]int8, m)

	type nb struct{ to, eid int }
	g := make([][]nb, n+1)
	dfn := make([]int, n+1)
	var f func(v, fid int) int
	f = func(v, fid int) int {
		c++
		dfn[v] = c
		lowV := c
		for _, e := range g[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := f(w, e.eid)
				if lowW > dfn[v] {
					ans[e.eid] = 1
				}
				lowV = min(lowV, lowW)
			} else if e.eid != fid {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for _, es := range edges {
		if es == nil {
			continue
		}
		vs := []int{}
		for _, e := range es {
			if v, w, i := find(e.v), find(e.w), e.i; v != w {
				g[v] = append(g[v], nb{w, i})
				g[w] = append(g[w], nb{v, i})
				vs = append(vs, v, w)
			} else {
				ans[i] = 2
			}
		}
		for _, v := range vs {
			if dfn[v] == 0 {
				f(v, -1)
			}
		}
		for i := 0; i < len(vs); i += 2 {
			v, w := vs[i], vs[i+1]
			fa[find(v)] = find(w)
			g[v] = nil
			g[w] = nil
			dfn[v] = 0
			dfn[w] = 0
		}
	}

	s := [3]string{"at least one", "any", "none"}
	for _, t := range ans {
		Fprintln(out, s[t])
	}
}

//func main() { CF160D(os.Stdin, os.Stdout) }
