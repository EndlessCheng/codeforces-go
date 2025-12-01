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
	var n, m, sum, v int
	Fscan(in, &n, &m)
	st := m + n*5
	end := st + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}

	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for j := 0; j < 4; j++ {
			addEdge(m+i*5+j, m+i*5+j+1, v*j)
		}
		addEdge(m+i*5+4, end, v*4)
	}
	for i := 0; i < m; i++ {
		Fscan(in, &v)
		addEdge(st, i, v)
		sum += v
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			addEdge(i, m+j*5+v-1, math.MaxInt)
		}
	}

	d := make([]int, len(g))
	bfs := func() bool {
		for i := range d {
			d[i] = 0
		}
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
		for i := range iter {
			iter[i] = 0
		}
		maxFlow += dfs(st, math.MaxInt)
	}

	Fprintln(out, sum-maxFlow)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
func min(a, b int) int { if b < a { return b }; return a }
