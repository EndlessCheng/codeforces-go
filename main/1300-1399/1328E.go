package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1328E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, v, w, k int
	Fscan(in, &n, &q)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mx = 19
	pa := make([][mx]int, n)
	dep := make([]int, n)
	var dfs func(v, p, d int)
	dfs = func(v, p, d int) {
		pa[v][0] = p
		dep[v] = d
		for _, w := range g[v] {
			if w != p {
				dfs(w, v, d+1)
			}
		}
	}
	dfs(0, -1, 0)
	for k := 0; k+1 < mx; k++ {
		for v := range pa {
			if p := pa[v][k]; p != -1 {
				pa[v][k+1] = pa[p][k]
			} else {
				pa[v][k+1] = -1
			}
		}
	}
	_lca := func(v, w int) int {
		if dep[v] > dep[w] {
			v, w = w, v
		}
		for k := 0; k < mx; k++ {
			if (dep[w]-dep[v])>>k&1 == 1 {
				w = pa[w][k]
			}
		}
		if v == w {
			return v
		}
		for k := mx - 1; k >= 0; k-- {
			if pa[v][k] != pa[w][k] {
				v, w = pa[v][k], pa[w][k]
			}
		}
		return pa[v][0]
	}

o:
	for ; q > 0; q-- {
		Fscan(in, &k)
		vs := make([]int, k)
		maxDep, maxV := 0, 0
		for i := range vs {
			Fscan(in, &v)
			v--
			if dep[v] > maxDep {
				maxDep, maxV = dep[v], v
			}
			vs[i] = v
		}
		for _, v := range vs {
			if dep[v]-dep[_lca(maxV, v)] > 1 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1328E(os.Stdin, os.Stdout) }
