package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1027D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, s int
	Fscan(in, &n)
	c := make([]int, n)
	for i := range c {
		Fscan(in, &c[i])
	}
	g := make([]int, n)
	deg := make([]int, n)
	for i := range g {
		Fscan(in, &g[i])
		g[i]--
		deg[g[i]]++
	}

	vis := make([]bool, n)
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		vis[v] = true
		w := g[v]
		deg[w]--
		if deg[w] == 0 {
			q = append(q, w)
		}
	}

	for i, b := range vis {
		if !b {
			min := int(1e9)
			for v := i; !vis[v]; v = g[v] {
				vis[v] = true
				if c[v] < min {
					min = c[v]
				}
			}
			s += min
		}
	}
	Fprint(out, s)
}

//func main() { CF1027D(os.Stdin, os.Stdout) }
