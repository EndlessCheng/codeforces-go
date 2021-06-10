package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1220E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, s int
	Fscan(in, &n, &m)
	a := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &a[i])
	}
	type edge struct{ v, w int }
	e := make([]edge, m)
	g := make([][]int, n)
	d := make([]int, n)
	for i := range e {
		Fscan(in, &v, &w)
		v--
		w--
		e[i] = edge{v, w}
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		d[v]++
		d[w]++
	}
	Fscan(in, &s)
	s--
	q := []int{}
	for i, d := range d {
		if d == 1 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v, q = q[0], q[1:]
		for _, w := range g[v] {
			if d[w]--; d[w] == 1 {
				q = append(q, w)
			}
		}
	}

	for i, d := range d {
		if d > 1 {
			a[n] += a[i]
		}
	}
	g2 := make([][]int, n+1)
	for _, e := range e {
		v, w := e.v, e.w
		if d[v] > 1 {
			v = n
		}
		if d[w] > 1 {
			w = n
		}
		if v != w {
			g2[v] = append(g2[v], w)
			g2[w] = append(g2[w], v)
		}
	}

	ans := int64(0)
	vis := make([]bool, n+1)
	var f func(int, int, int64)
	f = func(v, fa int, s int64) {
		vis[v] = true
		if s > ans {
			ans = s
		}
		if v == n {
			fa = -1
		}
		for _, w := range g2[v] {
			if w != fa {
				s := s
				if !vis[w] {
					s += a[w]
				}
				f(w, v, s)
			}
		}
	}
	if d[s] > 1 {
		s = n
	}
	f(s, -1, a[s])
	Fprint(out, ans)
}

//func main() { CF1220E(os.Stdin, os.Stdout) }
