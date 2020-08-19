package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF546E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var N, m, v, w, s, maxFlow int
	Fscan(in, &N, &m)
	st, end, n := 2*N, 2*N+1, 2*N+2
	type nb struct{ to, rid, cap int }
	g := make([][]nb, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}
	for i := 0; i < N; i++ {
		Fscan(in, &v)
		addEdge(st, i, v)
		s += v
	}
	sum := s
	for i := N; i < 2*N; i++ {
		Fscan(in, &v)
		addEdge(i, end, v)
		s -= v
	}
	if s != 0 {
		Fprint(out, "NO")
		return
	}
	for i := 0; i < N; i++ {
		addEdge(i, i+N, inf)
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		addEdge(v, w+N, inf)
		addEdge(w, v+N, inf)
	}

	dep := make([]int, n)
	bfs := func() bool {
		for i := range dep {
			dep[i] = -1
		}
		dep[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && dep[w] < 0 {
					dep[w] = dep[v] + 1
					q = append(q, w)
				}
			}
		}
		return dep[end] >= 0
	}
	var it []int
	var dfs func(int, int) int
	dfs = func(v, minF int) int {
		if v == end {
			return minF
		}
		for ; it[v] < len(g[v]); it[v]++ {
			e := &g[v][it[v]]
			if w := e.to; e.cap > 0 && dep[w] > dep[v] {
				if f := dfs(w, min(minF, e.cap)); f > 0 {
					e.cap -= f
					g[w][e.rid].cap += f
					return f
				}
			}
		}
		return 0
	}
	for bfs() {
		it = make([]int, n)
		for {
			if f := dfs(st, inf); f > 0 {
				maxFlow += f
			} else {
				break
			}
		}
	}
	if maxFlow < sum {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	for _, es := range g[:N] {
		ans := make([]int, N)
		for _, e := range es {
			if w := e.to; N <= w && w < 2*N {
				ans[w-N] = g[w][e.rid].cap
			}
		}
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF546E(os.Stdin, os.Stdout) }
