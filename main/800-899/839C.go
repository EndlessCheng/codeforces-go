package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF839C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v, w int
	Fscan(in, &n)
	g := make([][]int, n)
	g[0] = append(g[0], -1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	var f func(v, fa int) float64
	f = func(v, fa int) (l float64) {
		if len(g[v]) == 1 {
			return
		}
		p := 1 / float64(len(g[v])-1)
		for _, w := range g[v] {
			if w != fa {
				l += p * (f(w, v) + 1)
			}
		}
		return
	}
	Fprintf(out, "%.15f", f(0, -1))
}

//func main() { CF839C(os.Stdin, os.Stdout) }
