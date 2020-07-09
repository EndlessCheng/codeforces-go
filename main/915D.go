package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF915D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w int
	Fscan(in, &n, &m)
	g := make([][]int, n)
	dd := make([]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		dd[w]++
	}

	for i, d := range dd {
		if d == 0 {
			continue
		}
		deg := make([]int, n)
		copy(deg, dd)
		deg[i]--
		q := []int{}
		for v, d := range deg {
			if d == 0 {
				q = append(q, v)
			}
		}
		c := 0
		for len(q) > 0 {
			v, q = q[0], q[1:]
			c++
			for _, w := range g[v] {
				deg[w]--
				if deg[w] == 0 {
					q = append(q, w)
				}
			}
		}
		if c == n {
			Fprint(out, "YES")
			return
		}
	}
	Fprint(out, "NO")
}

//func main() { CF915D(os.Stdin, os.Stdout) }
