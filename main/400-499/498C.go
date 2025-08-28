package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf498C(in io.Reader, out io.Writer) {
	var n, m, v int
	Fscan(in, &n, &m)
	type pair struct{ p, e int }
	ps := make([][]pair, n)
	sum := make([]int, n+1)
	for i := range ps {
		Fscan(in, &v)
		for p := 2; p*p <= v; p++ {
			if v%p > 0 {
				continue
			}
			e := 1
			for v /= p; v%p == 0; v /= p {
				e++
			}
			ps[i] = append(ps[i], pair{p, e})
		}
		if v > 1 {
			ps[i] = append(ps[i], pair{v, 1})
		}
		sum[i+1] = sum[i] + len(ps[i])
	}

	st := sum[n]
	end := st + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}

	for i, ps := range ps {
		if i%2 == 0 {
			for j, p := range ps {
				addEdge(st, sum[i]+j, p.e)
			}
		} else {
			for j, p := range ps {
				addEdge(sum[i]+j, end, p.e)
			}
		}
	}

	for range m {
		var a, b int
		Fscan(in, &a, &b)
		a--
		b--
		if a%2 > 0 {
			a, b = b, a
		}
		for i, p := range ps[a] {
			for j, q := range ps[b] {
				if p.p == q.p {
					addEdge(sum[a]+i, sum[b]+j, 1e9)
				}
			}
		}
	}

	d := make([]int, len(g))
	bfs := func() bool {
		clear(d)
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
	iter := make([]int, len(g))
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == end {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && d[w] > d[v] {
				f := dfs(w, min(totalFlow-curFlow, e.cap))
				if f == 0 {
					continue
				}
				e.cap -= f
				g[w][e.rid].cap += f
				curFlow += f
				if curFlow == totalFlow {
					break
				}
			}
		}
		return
	}
	maxFlow := 0
	for bfs() {
		clear(iter)
		maxFlow += dfs(st, 1e9)
	}
	Fprint(out, maxFlow)
}

//func main() { cf498C(bufio.NewReader(os.Stdin), os.Stdout) }
