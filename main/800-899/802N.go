package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf802N(in io.Reader, out io.Writer) {
	const inf int = 1e18
	var n, k, v int
	Fscan(in, &n, &k)

	st, st2, end := n, n+1, n+2
	type nb struct{ to, rid, cap, cost int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
	}

	addEdge(st, st2, k, 0)
	for i := range n {
		Fscan(in, &v)
		addEdge(st2, i, 1, v)
		if i > 0 {
			addEdge(i-1, i, inf, 0) // 让时间流动！
		}
	}
	for i := range n {
		Fscan(in, &v)
		addEdge(i, end, 1, v)
	}

	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = inf
		}
		dis[st] = 0
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
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		return dis[end] < inf
	}

	minCost := 0
	for spfa() {
		minF := inf
		for v := end; v != st; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := end; v != st; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[end] * minF
	}
	Fprint(out, minCost)
}

//func main() { cf802N(bufio.NewReader(os.Stdin), os.Stdout) }
