package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1656E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		a := make([]int, n)
		var f func(v, fa, c int)
		f = func(v, fa, c int) {
			a[v] = len(g[v]) * c
			for _, w := range g[v] {
				if w != fa {
					f(w, v, -c)
				}
			}
		}
		f(0, -1, 1)
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1656E(os.Stdin, os.Stdout) }
