package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF618D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	var x, y int64
	Fscan(in, &n, &x, &y)
	g := make([][]int, n)
	star := false
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
		if len(g[v]) == n-1 || len(g[w]) == n-1 {
			star = true
		}
	}

	if x < y {
		pathEdge := 0
		var f func(int, int) bool
		f = func(v, fa int) bool {
			c := 0
			for _, w := range g[v] {
				if w != fa && f(w, v) {
					c++
				}
			}
			if c < 2 {
				pathEdge += c
				return true
			}
			pathEdge += 2
			return false
		}
		f(0, -1)
		Fprint(out, int64(pathEdge)*x+int64(n-1-pathEdge)*y)
	} else if star {
		Fprint(out, int64(n-2)*y+x)
	} else {
		Fprint(out, int64(n-1)*y)
	}
}

//func main() { CF618D(os.Stdin, os.Stdout) }
