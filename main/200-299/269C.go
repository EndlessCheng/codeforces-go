package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF269C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v, w, c int
	Fscan(in, &n, &m)
	type edge struct{ to, i, d, f int }
	g := make([][]edge, n+1)
	f := make([]int, n+1)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w, &c)
		g[v] = append(g[v], edge{w, i, 0, c})
		g[w] = append(g[w], edge{v, i, 1, c})
		f[v] += c
		f[w] += c
	}
	for i := range f {
		f[i] >>= 1
	}

	ans := make([]int, m)
	for i := range ans {
		ans[i] = -1
	}
	q := []int{1}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range g[v] {
			if ans[e.i] == -1 {
				ans[e.i] = e.d
				w := e.to
				f[w] -= e.f
				if w != n && f[w] == 0 {
					q = append(q, w)
				}
			}
		}
	}
	for _, d := range ans {
		Fprintln(out, d)
	}
}

//func main() { CF269C(os.Stdin, os.Stdout) }
