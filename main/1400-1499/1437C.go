package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1437C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n, t int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		type nb struct{ to, rid, cap, cost int }
		g := make([][]nb, 3*n+2)
		addEdge := func(from, to, cap, cost int) {
			g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
			g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
		}
		st, end := 0, 3*n+1
		for i := 1; i <= n; i++ {
			addEdge(st, i, 1, 0)
		}
		for i := n + 1; i <= 3*n; i++ {
			addEdge(i, end, 1, 0)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &t)
			for j := 1; j <= 2*n; j++ {
				addEdge(i, n+j, inf, abs(t-j))
			}
		}

		n = 3*n + 2
		dist := make([]int, n)
		type pair struct{ v, i int }
		fa := make([]pair, n)
		spfa := func() bool {
			for i := range dist {
				dist[i] = inf
			}
			dist[st] = 0
			inQ := make([]bool, n)
			inQ[st] = true
			q := []int{st}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				inQ[v] = false
				for i, e := range g[v] {
					if e.cap == 0 {
						continue
					}
					w := e.to
					if newD := dist[v] + e.cost; newD < dist[w] {
						dist[w] = newD
						fa[w] = pair{v, i}
						if !inQ[w] {
							q = append(q, w)
							inQ[w] = true
						}
					}
				}
			}
			return dist[end] < inf
		}
		minCost := 0
		for spfa() {
			for v := end; v != st; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap--
				g[v][e.rid].cap++
				v = p.v
			}
			minCost += dist[end]
		}
		Fprintln(out, minCost)
	}
}

//func main() { CF1437C(os.Stdin, os.Stdout) }
