package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	m := n
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			Fscan(in, &a[i][j])
		}
	}

	type nb struct{ to, rid, cap, cost int }
	var g [][]nb
	f := func(lim int) (tarF, minCost int) {
		S := n + m
		T := S + 1
		g = make([][]nb, T+1, T+2)
		addEdge := func(from, to, cap, cost int) {
			g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
			g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
		}
		for i, row := range a {
			for j, v := range row {
				addEdge(i, n+j, 1, -v)
			}
			addEdge(S, i, k, 0)
		}
		for j := range a[0] {
			addEdge(n+j, T, k, 0)
		}
		if lim >= 0 {
			g = append(g, []nb{})
			addEdge(T+1, S, lim, 0)
			S = T + 1
		}

		dis := make([]int, len(g))
		type vi struct{ v, i int }
		fa := make([]vi, len(g))
		inQ := make([]bool, len(g))
		spfa := func() bool {
			for i := range dis {
				dis[i] = math.MaxInt
			}
			dis[S] = 0
			inQ[S] = true
			q := []int{S}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				inQ[v] = false
				for i, e := range g[v] {
					if e.cap == 0 {
						continue
					}
					w := e.to
					newD := dis[v] + e.cost
					if newD < dis[w] {
						dis[w] = newD
						fa[w] = vi{v, i}
						if !inQ[w] {
							inQ[w] = true
							q = append(q, w)
						}
					}
				}
			}
			return dis[T] < math.MaxInt
		}

		maxFlow := 0
		minC := math.MaxInt
		for spfa() {
			minF := math.MaxInt
			for v := T; v != S; {
				p := fa[v]
				minF = min(minF, g[p.v][p.i].cap)
				v = p.v
			}
			for v := T; v != S; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dis[T] * minF
			if minCost < minC {
				minC = minCost
				tarF = maxFlow
			}
		}
		return
	}
	tarF, _ := f(-1)
	_, cost := f(tarF)
	Fprintln(out, -cost)
	for _, to := range g[:n] {
		for _, e := range to[:m] {
			if e.cap == 0 {
				Fprint(out, "X")
			} else {
				Fprint(out, ".")
			}
		}
		Fprintln(out)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
