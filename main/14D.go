package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF14D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, ans int
	Fscan(in, &n)
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	d := func(u, no int) (maxD int) {
		var f func(v, fa, d int)
		f = func(v, fa, d int) {
			if d > maxD {
				maxD = d
				u = v
			}
			for _, w := range g[v] {
				if w != fa && w != no {
					f(w, v, d+1)
				}
			}
		}
		maxD = -1
		f(u, -1, 0)
		maxD = -1
		f(u, -1, 0)
		return
	}
	var f func(v, fa int)
	f = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				if s := d(v, w) * d(w, v); s > ans {
					ans = s
				}
				f(w, v)
			}
		}
	}
	f(0, -1)
	Fprint(out, ans)
}

//func main() { CF14D(os.Stdin, os.Stdout) }
