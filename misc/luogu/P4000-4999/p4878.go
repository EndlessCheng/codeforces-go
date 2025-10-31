package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func p4878(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e9

	var n, m1, m2, v, w, wt int
	Fscan(in, &n, &m1, &m2)
	type nb struct{ to, wt int }
	g := make([][]nb, n+1)
	for i := 1; i <= n; i++ {
		g[0] = append(g[0], nb{i, 0})
	}
	for ; m1 > 0; m1-- {
		Fscan(in, &v, &w, &wt)
		g[v] = append(g[v], nb{w, wt})
	}
	for ; m2 > 0; m2-- {
		Fscan(in, &v, &w, &wt)
		g[w] = append(g[w], nb{v, -wt})
	}

	spfa := func(st int) int {
		dist := make([]int, n+1)
		for i := range dist {
			dist[i] = inf
		}
		dist[st] = 0
		inQ := make([]bool, n+1)
		inQ[st] = true
		relaxedCnt := make([]int, n+1)
		q := []int{st}
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
						return -1
					}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		if dist[n] == inf {
			return -2
		}
		return dist[n]
	}
	d := spfa(0)
	if d >= 0 {
		d = spfa(1)
	}
	Fprint(out, d)
}

//func main() { p4878(os.Stdin, os.Stdout) }
