package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1082G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int = 1e18
	var n, m, cost, v, w, profit, tot int
	Fscan(in, &n, &m)

	st := 0
	end := n + m + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &cost)
		addEdge(m+i, end, cost)
	}
	for i := 1; i <= m; i++ {
		Fscan(in, &v, &w, &profit)
		addEdge(st, i, profit)
		addEdge(i, m+v, inf)
		addEdge(i, m+w, inf)
		tot += profit
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
			w := e.to
			if e.cap > 0 && d[w] > d[v] {
				f := dfs(w, min(minF, e.cap))
				if f > 0 {
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
				f := dfs(st, inf)
				if f > 0 {
					maxFlow += f
				} else {
					break
				}
			}
		}
		return
	}
	Fprint(out, tot-dinic())
}

//func main() { cf1082G(os.Stdin, os.Stdout) }
