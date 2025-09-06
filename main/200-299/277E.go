package main

import (
	. "fmt"
	"io"
	"math"
)

/*
流量
可以用来“选择”或“计数”。我们可以用 1 单位的流量代表“选择一条边”。
因此，要构成一棵树，我们总共需要选择 n-1 条边，这意味着网络中的总流量应该是 n-1。

容量
可以用来施加“约束”。例如，一个节点的出度不能超过 2，就可以通过设置相关边的容量来实现。

费用
可以用来衡量“成本”。边的长度自然就是我们想要最小化的成本。
*/

// https://github.com/EndlessCheng
func cf277E(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]struct{ x, y int }, n)
	for i := range a {
		Fscan(in, &a[i].x, &a[i].y)
	}

	S := n * 2
	T := S + 1
	type nb struct {
		to, rid, cap int
		cost         float64
	}
	g := make([][]nb, T+1)
	addEdge := func(from, to, cap int, cost float64) {
		g[from] = append(g[from], nb{to, len(g[to]), cap, cost})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0, -cost})
	}
	for i, p := range a {
		addEdge(S, i, 2, 0)
		addEdge(n+i, T, 1, 0)
		for j, q := range a {
			if p.y > q.y {
				addEdge(i, n+j, 1, math.Sqrt(float64((p.x-q.x)*(p.x-q.x)+(p.y-q.y)*(p.y-q.y))))
			}
		}
	}

	dis := make([]float64, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxFloat64
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
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
		return dis[T] < math.MaxFloat64
	}
	maxFlow := 0
	minCost := 0.
	for spfa() {
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		maxFlow += minF
		minCost += dis[T] * float64(minF)
	}

	if maxFlow == n-1 {
		Fprintf(out, "%.6f", minCost)
	} else {
		Fprint(out, -1)
	}
}

//func main() { cf277E(os.Stdin, os.Stdout) }
