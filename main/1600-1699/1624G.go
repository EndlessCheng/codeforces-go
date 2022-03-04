package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1624G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w, wt int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type nb struct{ to, wt int }
		g := make([][]nb, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w, &wt)
			v--
			w--
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
		}
		vis := make([]int8, n)
		ans := 0
		for i, m := int8(29), 0; i >= 0; i-- {
			c := n
			m |= 1 << i
			var f func(int)
			f = func(v int) {
				vis[v] = i + 1
				c--
				for _, e := range g[v] {
					if e.wt&m == 0 && vis[e.to] != i+1 {
						f(e.to)
					}
				}
			}
			f(0)
			if c > 0 {
				ans |= 1 << i
				m &^= 1 << i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1624G(os.Stdin, os.Stdout) }
