package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF963B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, w int
	Fscan(in, &n)
	if n&1 == 0 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	g := make([][]int, n+1)
	for v := 1; v <= n; v++ {
		if Fscan(in, &w); w > 0 {
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
	}
	size := make([]int, n+1)
	var f func(v, fa int) int
	f = func(v, fa int) int {
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				sz += f(w, v)
			}
		}
		size[v] = sz
		return sz
	}
	f(1, 0)
	var f2 func(v, fa int)
	f2 = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa && size[w]&1 == 0 {
				f2(w, v)
			}
		}
		Fprintln(out, v)
		for _, w := range g[v] {
			if w != fa && size[w]&1 > 0 {
				f2(w, v)
			}
		}
	}
	f2(1, 0)
}

//func main() { CF963B(os.Stdin, os.Stdout) }
