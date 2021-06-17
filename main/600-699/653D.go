package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF653D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	binarySearchF := func(l, r float64, f func(x float64) bool) float64 {
		// 由于范围为 1e6，精度为 1e-8，且要乘上一个 1e5 的整数，循环次数设置为 log2(1e19)
		for step := 63; step > 0; step-- {
			m := (l + r) / 2
			if f(m) {
				r = m
			} else {
				l = m
			}
		}
		return (l + r) / 2
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, m, b, v, w, cap, st int
	Fscan(in, &n, &m, &b)
	end := n - 1
	type nb struct{ to, rid, cap, c int }
	g := make([][]nb, n)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap, 0})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, 0})
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w, &cap)
		v--
		w--
		addEdge(v, w, cap)
	}

	d := make([]int, n)
	bfs := func() bool {
		for i := range d {
			d[i] = -1
		}
		d[st] = 0
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.c > 0 && d[w] < 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d[end] >= 0
	}
	var iter []int
	var dfs func(int, int) int
	dfs = func(v int, minF int) int {
		if v == end {
			return minF
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.c > 0 && d[w] > d[v] {
				if f := dfs(w, min(minF, e.c)); f > 0 {
					e.c -= f
					g[w][e.rid].c += f
					return f
				}
			}
		}
		return 0
	}

	ans := binarySearchF(0, 1e6, func(x float64) bool {
		for _, es := range g {
			for j := range es {
				c := int64(float64(es[j].cap) / x) // 可能爆 int32
				if c > int64(b) {
					c = int64(b)
				}
				es[j].c = int(c)
			}
		}
		maxFlow := 0
		for bfs() {
			iter = make([]int, n)
			for {
				if f := dfs(st, 1e9); f > 0 {
					maxFlow += f
				} else {
					break
				}
			}
		}
		return maxFlow < b
	})
	Fprintf(out, "%.10f", ans*float64(b))
}

//func main() { CF653D(os.Stdin, os.Stdout) }
