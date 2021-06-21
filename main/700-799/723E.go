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
	type edge struct{ to, i int }

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]edge, n+1)
		for i := 0; i < m; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], edge{w, i})
			g[w] = append(g[w], edge{v, i})
		}
		ans := -1
		for i, vs := range g {
			if len(vs)&1 > 0 {
				g[0] = append(g[0], edge{i, m})
				g[i] = append(g[i], edge{0, m})
				m++
			} else {
				ans++
			}
		}
		Fprintln(out, ans)

		pre := -1
		visV := make([]bool, n+1)
		visE := make([]bool, m)
		var f func(int)
		f = func(v int) {
			visV[v] = true
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				i := e.i
				if visE[i] {
					continue
				}
				visE[i] = true
				f(e.to)
				v, w := v, e.to
				if w == pre {
					v, w = w, v
				}
				if v > 0 && w > 0 {
					Fprintln(out, v, w)
				}
				pre = w
			}
		}
		for i, b := range visV {
			if !b {
				pre = -1
				f(i)
			}
		}
	}
}

//func main() { CF723E(os.Stdin, os.Stdout) }
