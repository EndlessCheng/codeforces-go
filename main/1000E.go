package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 另一种写法，不需要缩点 https://codeforces.com/problemset/submission/1000/83364650

// github.com/EndlessCheng/codeforces-go
func CF1000E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type neighbor struct{ to, eid int }
	type edge struct{ v, w int }

	var n, m, v, w, u, ans int
	Fscan(in, &n, &m)
	g := make([][]neighbor, n)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], neighbor{w, i})
		g[w] = append(g[w], neighbor{v, i})
		edges[i] = edge{v, w}
	}

	isBridge := make([]bool, m)
	dfn := make([]int, n)
	dfsClock := 0
	var f func(v, fa int) int
	f = func(v, fa int) int {
		dfsClock++
		dfn[v] = dfsClock
		lowV := dfsClock
		for _, e := range g[v] {
			w := e.to
			if dfn[w] == 0 {
				lowW := f(w, v)
				if lowW > dfn[v] {
					isBridge[e.eid] = true
				}
				lowV = min(lowV, lowW)
			} else if w != fa && dfn[w] < dfn[v] {
				lowV = min(lowV, dfn[w])
			}
		}
		return lowV
	}
	for i, t := range dfn {
		if t == 0 {
			f(i, -1)
		}
	}
	ids := make([]int, n)
	idCnt := 0
	var f2 func(int)
	f2 = func(v int) {
		ids[v] = idCnt
		for _, e := range g[v] {
			if w := e.to; ids[w] == 0 && !isBridge[e.eid] {
				f2(w)
			}
		}
	}
	for i, id := range ids {
		if id == 0 {
			idCnt++
			f2(i)
		}
	}

	t := make([][]int, idCnt+1)
	for _, e := range edges {
		if v, w := ids[e.v], ids[e.w]; v != w {
			t[v] = append(t[v], w)
			t[w] = append(t[w], v)
		}
	}
	var f3 func(v, fa, d int)
	f3 = func(v, fa, d int) {
		if d > ans {
			ans = d
			u = v
		}
		for _, w := range t[v] {
			if w != fa {
				f3(w, v, d+1)
			}
		}
	}
	ans = -1
	f3(1, 0, 0)
	ans = -1
	f3(u, 0, 0)
	Fprint(out, ans)
}

//func main() { CF1000E(os.Stdin, os.Stdout) }
