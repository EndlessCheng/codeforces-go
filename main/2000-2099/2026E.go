package main

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2026E(in io.Reader, out io.Writer) {
	var T, n int
	var s uint
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		const mx = 60
		st := n + mx
		end := st + 1
		type nb struct{ to, rid, cap int }
		g := make([][]nb, end+1)
		addEdge := func(from, to, cap int) {
			g[from] = append(g[from], nb{to, len(g[to]), cap})
			g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
		}
		for i := range n {
			for Fscan(in, &s); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				addEdge(i, n+j, 1)
			}
			addEdge(st, i, 1)
		}
		for j := range mx {
			addEdge(n+j, end, 1)
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
			maxFlow += dfs(st, math.MaxInt)
		}

		Fprintln(out, n-maxFlow)
	}
}

//func main() { cf2026E(bufio.NewReader(os.Stdin), os.Stdout) }
