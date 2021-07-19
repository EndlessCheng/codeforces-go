package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF734E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, u, mxD int
	Fscan(in, &n)
	c := make([]int8, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		if d > mxD {
			mxD, u = d, v
		}
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+int(c[v]^c[w]))
			}
		}
	}
	f(0, -1, 1)
	mxD = 0
	f(u, -1, 1)
	Fprint(out, mxD/2)
}

//func main() { CF734E(os.Stdin, os.Stdout) }
