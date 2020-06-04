package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF25D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type edge struct{ v, w int }
	deletes := []edge{}
	vis := make([]bool, n)
	var f func(v, fa int)
	f = func(v, fa int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w, v)
			} else if w != fa && v < w {
				deletes = append(deletes, edge{v, w})
			}
		}
	}
	vs := []int{}
	for i, b := range vis {
		if !b {
			vs = append(vs, i)
			f(i, -1)
		}
	}

	Fprintln(out, len(deletes))
	for i, p := range deletes {
		Fprintln(out, p.v+1, p.w+1, vs[i]+1, vs[i+1]+1)
	}
}

//func main() { CF25D(os.Stdin, os.Stdout) }
