package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func runB(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9

	var T, n, x, y int
	var tmp float64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &tmp)
		lim := int64(tmp * tmp)

		st := n * 2
		end := 0
		type nb struct {
			to, rid, cap int
			forw         bool
		}
		g := make([][]nb, st+1)
		addEdge := func(from, to, cap int) {
			g[from] = append(g[from], nb{to, len(g[to]), cap, true})
			g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, false})
		}
		targetFlow := 0
		a := make([]struct{ x, y int64 }, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y, &x, &y)
			addEdge(st, i, x)
			addEdge(i, i+n, y)
			targetFlow += x
		}
		for i, p := range a {
			for j, q := range a[:i] {
				if (p.x-q.x)*(p.x-q.x)+(p.y-q.y)*(p.y-q.y) <= lim {
					addEdge(i+n, j, inf)
					addEdge(j+n, i, inf)
				}
			}
		}

		var d []int
		bfs := func() bool {
			d = make([]int, len(g))
			d[st] = 1
			q := []int{st}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				for _, e := range g[v] {
					if w := e.to; e.cap > 0 && d[w] == 0 {
						d[w] = d[v] + 1
						q = append(q, w)
					}
				}
			}
			return d[end] > 0
		}
		var iter []int
		var dfs func(int, int) int
		dfs = func(v, minF int) int {
			if v == end {
				return minF
			}
			for ; iter[v] < len(g[v]); iter[v]++ {
				e := &g[v][iter[v]]
				if w := e.to; e.cap > 0 && d[w] > d[v] {
					if f := dfs(w, min(minF, e.cap)); f > 0 {
						e.cap -= f
						g[w][e.rid].cap += f
						return f
					}
				}
			}
			return 0
		}
		dinic := func() (maxFlow int) {
			for bfs() {
				iter = make([]int, len(g))
				for {
					if f := dfs(st, inf); f > 0 {
						maxFlow += f
					} else {
						break
					}
				}
			}
			return
		}
		found := false
		for ; end < n; end++ {
			for _, es := range g {
				for i, e := range es {
					if e.forw {
						es[i].cap += g[e.to][e.rid].cap
						g[e.to][e.rid].cap = 0
					}
				}
			}
			if dinic() == targetFlow {
				found = true
				Fprint(out, end, " ")
			}
		}
		if !found {
			Fprint(out, -1)
		}
		Fprintln(out)
	}
}

func main() { runB(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
