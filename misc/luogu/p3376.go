package main

import (
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func p3376(in io.Reader, out io.Writer) {
	var n, m, st, end, maxFlow int
	Fscan(in, &n, &m, &st, &end)
	st--
	end--
	type nb struct{ to, rid, cap int }
	g := make([][]nb, n)
	for ; m > 0; m-- {
		var v, w, cap_ int
		Fscan(in, &v, &w, &cap_)
		v--
		w--
		g[v] = append(g[v], nb{w, len(g[w]), cap_})
		g[w] = append(g[w], nb{v, len(g[v]) - 1, 0})
	}

	d := make([]int, n)
	bfs := func() bool {
		clear(d) // d[i] = 0 表示没有访问过
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
	// 当前弧，在其之前的边已经没有用了，避免对没有用的边进行多次检查
	iter := make([]int, n)
	// 寻找增广路（多路增广）
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
	for bfs() {
		clear(iter)
		maxFlow += dfs(st, math.MaxInt)
	}
	Fprint(out, maxFlow)
}

//func main() { p3376(bufio.NewReader(os.Stdin), os.Stdout) }
