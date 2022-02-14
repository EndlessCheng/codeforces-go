package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF723E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type nb struct{ to, i int }
		g := make([][]nb, n+1)
		for i := 0; i < m; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], nb{w, i})
			g[w] = append(g[w], nb{v, i})
		}
		ans := -1
		for i, vs := range g {
			if len(vs)&1 > 0 {
				g[0] = append(g[0], nb{i, m})
				g[i] = append(g[i], nb{0, m})
				m++
			} else {
				ans++
			}
		}
		Fprintln(out, ans)

		vis := make([]bool, m)
		var f func(int)
		f = func(v int) {
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				i := e.i
				if vis[i] {
					continue
				}
				vis[i] = true
				w := e.to
				f(w)
				if v > 0 && w > 0 {
					Fprintln(out, v, w)
				}
			}
		}
		for i := range g {
			f(i)
		}
	}
}

//func main() { CF723E(os.Stdin, os.Stdout) }
