package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF500D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ l, c int64 }
	type edge struct{ to, id int }

	var n, v, w, q, id int
	Fscan(in, &n)
	es := make([]pair, n)
	g := make([][]edge, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &es[i].l)
		v--
		w--
		g[v] = append(g[v], edge{w, i})
		g[w] = append(g[w], edge{v, i})
	}

	size := make([]int, n)
	var _f func(v, fa int) int
	_f = func(v, fa int) int {
		sz := 1
		for _, e := range g[v] {
			if w := e.to; w != fa {
				sz += _f(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	_f(0, -1)

	var f func(v, fa int)
	f = func(v, fa int) {
		for _, e := range g[v] {
			if w := e.to; w != fa {
				if s, z := int64(size[w]), int64(n-size[w]); s > 1 && z > 1 {
					es[e.id].c = s * z * (s + z - 2)
				} else {
					es[e.id].c = int64(n-1) * int64(n-2)
				}
				f(w, v)
			}
		}
	}
	f(0, -1)
	s := 0.
	for _, p := range es[1:] {
		s += float64(p.l * p.c)
	}
	c := float64(int64(n) * int64(n-1) * int64(n-2) / 6)
	for Fscan(in, &q); q > 0; q-- {
		var l int64
		Fscan(in, &id, &l)
		e := es[id]
		s -= float64(e.c * (e.l - l))
		Fprintf(out, "%.7f\n", s/c)
		es[id].l = l
	}
}

//func main() { CF500D(os.Stdin, os.Stdout) }
