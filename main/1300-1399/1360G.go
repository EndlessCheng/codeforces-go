package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1360G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	type edge struct{ to, rev, cap int }

	solve := func(_case int) {
		var n, m, a, b, maxFlow int
		Fscan(in, &n, &m, &a, &b)
		if n*a != m*b {
			Fprintln(out, "NO")
			return
		}

		edges := make([][]edge, n+m+2)
		addEdge := func(from, to int, cap int) {
			edges[from] = append(edges[from], edge{to, len(edges[to]), cap})
			edges[to] = append(edges[to], edge{from, len(edges[from]) - 1, 0})
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				addEdge(i, n+j, 1)
			}
		}
		st, end := n+m, n+m+1
		for i := 0; i < n; i++ {
			addEdge(st, i, a)
		}
		for i := 0; i < m; i++ {
			addEdge(n+i, end, b)
		}

		level := make([]int, n+m+2)
		calcLevel := func() bool {
			for i := range level {
				level[i] = -1
			}
			level[st] = 0
			q := []int{st}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				for _, e := range edges[v] {
					if w := e.to; e.cap > 0 && level[w] < 0 {
						level[w] = level[v] + 1
						q = append(q, w)
					}
				}
			}
			return level[end] >= 0
		}
		var iter []int
		var dfs func(int, int) int
		dfs = func(v int, mf int) int {
			if v == end {
				return mf
			}
			for i := iter[v]; i < len(edges[v]); i++ {
				e := &edges[v][i]
				if w := e.to; e.cap > 0 && level[w] > level[v] {
					if f := dfs(w, min(mf, e.cap)); f > 0 {
						e.cap -= f
						edges[w][e.rev].cap += f
						return f
					}
				}
				iter[v]++
			}
			return 0
		}
		const inf int = 1e9
		for calcLevel() {
			iter = make([]int, n+m+2)
			for {
				if f := dfs(st, inf); f > 0 {
					maxFlow += f
				} else {
					break
				}
			}
		}
		if maxFlow < n*a {
			Fprintln(out, "NO")
			return
		}
		Fprintln(out, "YES")
		for i := 0; i < n; i++ {
			ans := make([]byte, m)
			for j, e := range edges[i][:m] {
				ans[j] = '1' - byte(e.cap)
			}
			Fprintln(out, string(ans))
		}
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		solve(_case)
	}
}

//func main() { CF1360G(os.Stdin, os.Stdout) }
