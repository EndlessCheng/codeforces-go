package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1404B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, a, b, da, db, dab, v, w, u, maxD int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &a, &b, &da, &db)
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			if v == b {
				dab = d
			}
			if d > maxD {
				maxD, u = d, v
			}
			for _, w := range g[v] {
				if w != fa {
					f(w, v, d+1)
				}
			}
		}
		maxD = -1
		f(a, 0, 0)
		if 2*da >= db || da >= dab {
			Fprintln(out, "Alice")
			continue
		}
		maxD = -1
		f(u, 0, 0)
		if 2*da >= maxD {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { CF1404B(os.Stdin, os.Stdout) }
