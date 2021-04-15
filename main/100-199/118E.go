package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF118E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, v, w int
	Fscan(in, &n, &m)
	type nb struct{ to, i int }
	g := make([][]nb, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], nb{w, i})
		g[w] = append(g[w], nb{v, i})
	}
	dfn := make([]int, n+1)
	ts := 0
	var f func(int, int) int
	f = func(v, fa int) int {
		ts++
		dfn[v] = ts
		lowV := ts
		for _, e := range g[v] {
			if w := e.to; dfn[w] == 0 {
				lowW := f(w, v)
				if lowW > dfn[v] {
					return 1e9
				}
				lowV = min(lowV, lowW)
			} else if w != fa {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	if f(1, 0) == 1e9 {
		Fprint(out, 0)
		return
	}
	vis := make([]bool, m)
	var f2 func(int)
	f2 = func(v int) {
		for len(g[v]) > 0 {
			e := g[v][0]
			g[v] = g[v][1:]
			if !vis[e.i] {
				vis[e.i] = true
				Fprintln(out, v, e.to)
				f2(e.to)
			}
		}
	}
	f2(1)
}

//func main() { CF118E(os.Stdin, os.Stdout) }
