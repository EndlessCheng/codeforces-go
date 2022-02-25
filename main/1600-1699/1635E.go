package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1635E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	es := make([][3]int, m)
	g := make([][]int, n)
	for i := range es {
		Fscan(in, &es[i][0], &es[i][1], &es[i][2])
		v, w := es[i][1]-1, es[i][2]-1
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	col := make([]int8, n)
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		col[v] = c
		for _, w := range g[v] {
			if col[w] == c || col[w] == 0 && !f(w, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range col {
		if c == 0 && !f(i, 1) {
			Fprint(out, "NO")
			return
		}
	}

	g = make([][]int, n)
	d := make([]int, n)
	for _, e := range es {
		v, w := e[1]-1, e[2]-1
		if col[v] == 2 {
			v, w = w, v
		}
		if e[0] == 1 {
			g[v] = append(g[v], w)
			d[w]++
		} else {
			g[w] = append(g[w], v)
			d[v]++
		}
	}
	q := make([]int, 0, n)
	for i, d := range d {
		if d == 0 {
			q = append(q, i)
		}
	}
	ans := make([]int, n)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		ans[v] = cap(q)
		for _, w := range g[v] {
			if d[w]--; d[w] == 0 {
				q = append(q, w)
			}
		}
	}
	if cap(q) > 0 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for i, v := range ans {
		Fprintf(out, "%c %d\n", " RL"[col[i]], v)
	}
}

//func main() { CF1635E(os.Stdin, os.Stdout) }
