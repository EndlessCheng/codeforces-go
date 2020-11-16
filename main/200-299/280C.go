package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF280C(_r io.Reader, _w io.Writer) {
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

	ans := .0
	var f func(v, fa, d int)
	f = func(v, fa, d int) {
		ans += 1 / float64(d)
		for _, w := range g[v] {
			if w != fa {
				f(w, v, d+1)
			}
		}
	}
	f(0, -1, 1)
	Fprintf(out, "%.8f", ans)
}

//func main() { CF280C(os.Stdin, os.Stdout) }
