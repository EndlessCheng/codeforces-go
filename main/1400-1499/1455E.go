package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1455E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int64 = 1e18
	perm4 := [][]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0},
	}

	var T int
	var x, y [4]int
	for Fscan(in, &T); T > 0; T-- {
		for i := 0; i < 4; i++ {
			Fscan(in, &x[i], &y[i])
		}
		ans := inf
		for _, p := range perm4 {
			type neighbor struct {
				to, rid   int
				cap, cost int64
			}
			g := [8][]neighbor{}
			addEdge := func(from, to int, cap, cost int64) {
				g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
				g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
			}
			for i := 0; i < 4; i++ {
				for j := 4; j < 6; j++ {
					addEdge(i, j, inf, 1)
					addEdge(j, i, inf, 1)
				}
			}
			st, end := 6, 7
			for i, w := range []int{
				x[p[2]] - x[p[0]],
				x[p[1]] - x[p[3]],
				y[p[0]] - y[p[1]],
				y[p[3]] - y[p[2]],
				x[p[0]] + y[p[2]] - x[p[1]] - y[p[0]],
				x[p[3]] + y[p[1]] - x[p[2]] - y[p[3]],
			} {
				if w > 0 {
					addEdge(st, i, int64(w), 0)
				} else {
					addEdge(i, end, int64(-w), 0)
				}
			}

			dist := [8]int64{}
			type pair struct{ v, i int }
			fa := [8]pair{}
			spfa := func() bool {
				for i := 0; i < 8; i++ {
					dist[i] = inf
				}
				dist[st] = 0
				inQ := [8]bool{}
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
			var maxFlow, minCost int64
			for spfa() {
				minF := inf
				for v := end; v != st; {
					p := fa[v]
					if c := g[p.v][p.i].cap; c < minF {
						minF = c
					}
					v = p.v
				}
				for v := end; v != st; {
					p := fa[v]
					e := &g[p.v][p.i]
					e.cap -= minF
					g[v][e.rid].cap += minF
					v = p.v
				}
				maxFlow += minF
				minCost += dist[end] * minF
			}
			if minCost < ans {
				ans = minCost
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1455E(os.Stdin, os.Stdout) }
