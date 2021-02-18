package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1487C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type nb struct{ to, i int }

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]nb, n)
		m := 0
		for i := 0; i < n; i++ {
			j := i + 1
			if n&1 == 0 && i&1 == 0 {
				j++
				m++
			}
			for ; j < n; j++ {
				g[i] = append(g[i], nb{j, m})
				g[j] = append(g[j], nb{i, m})
				m++
			}
		}

		ans := make([]int, m)
		vis := make([]bool, m)
		pre := 0
		var f func(int)
		f = func(v int) {
			for len(g[v]) > 0 {
				e := g[v][0]
				g[v] = g[v][1:]
				if w, i := e.to, e.i; !vis[i] {
					vis[i] = true
					f(w)
					v := v
					if w == pre {
						v, w = w, v
					}
					if v < w {
						ans[i] = 1
					} else {
						ans[i] = -1
					}
					pre = w
				}
			}
		}
		f(0)
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1487C(os.Stdin, os.Stdout) }
