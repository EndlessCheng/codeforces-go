package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func p2850(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t, n, m, m2, v, w, wt int
o:
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &m, &m2)
		type nb struct{ to, wt int }
		g := make([][]nb, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w, &wt)
			v--
			w--
			g[v] = append(g[v], nb{w, wt})
			g[w] = append(g[w], nb{v, wt})
		}
		for ; m2 > 0; m2-- {
			Fscan(in, &v, &w, &wt)
			g[v-1] = append(g[v-1], nb{w - 1, -wt})
		}
		dist := make([]int, n)
		q := make([]int, n)
		inQ := make([]bool, n)
		for i := 0; i < n; i++ {
			q[i] = i
			inQ[i] = true
		}
		relaxedCnt := make([]int, n)
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + e.wt; newD < dist[w] {
					dist[w] = newD
					relaxedCnt[w] = relaxedCnt[v] + 1
					if relaxedCnt[w] >= n {
						Fprintln(out, "YES")
						continue o
					}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		Fprintln(out, "NO")
	}
}

func main() { p2850(os.Stdin, os.Stdout) }
