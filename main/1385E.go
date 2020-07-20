package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1385E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ to, tp int }

	var t, n, m, tp, v, w int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m)
		g := make([][]pair, n)
		deg := make([]int, n)
		for ; m > 0; m-- {
			Fscan(in, &tp, &v, &w)
			v--
			w--
			g[v] = append(g[v], pair{w, tp})
			if tp == 0 {
				g[w] = append(g[w], pair{v, tp})
			} else {
				deg[w]++
			}
		}
		vis := make([]int8, n)
		q := []int{}
		for i, d := range deg {
			if d == 0 {
				vis[i] = 1
				q = append(q, i)
			}
		}
		c := 0
		for len(q) > 0 {
			v, q = q[0], q[1:]
			c++
			for i, e := range g[v] {
				w := e.to
				if e.tp == 0 {
					if vis[w] < 2 {
						g[v][i].tp = 1
					}
				} else {
					deg[w]--
				}
				if deg[w] == 0 && vis[w] == 0 {
					vis[w] = 1
					q = append(q, w)
				}
			}
			vis[v] = 2
		}
		if c < n {
			Fprintln(out, "NO")
			continue
		}
		Fprintln(out, "YES")
		for v, es := range g {
			for _, e := range es {
				if e.tp == 1 {
					Fprintln(out, v+1, e.to+1)
				}
			}
		}
	}
}

//func main() { CF1385E(os.Stdin, os.Stdout) }
