package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf311E(in io.Reader, out io.Writer) {
	var n, m, fg, k, tar, w, id, ans int
	Fscan(in, &n, &m, &fg)

	st := n + m
	end := st + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}

	sex := make([]bool, n)
	for i := range sex {
		Fscan(in, &sex[i])
	}
	for i, s := range sex {
		Fscan(in, &k)
		if !s {
			addEdge(st, i, k) // 如果不割 S->母狗，那么母狗就在 S 中
		} else {
			addEdge(i, end, k) // 如果不割公狗->T，那么公狗就在 T 中
		}
	}
	for i := range m {
		Fscan(in, &tar, &w, &k)
		for range k {
			Fscan(in, &id)
			id--
			if tar == 0 {
				addEdge(n+i, id, 1e18) // 如果不割 S->富人，那么母狗必须在 S 中
			} else {
				addEdge(id, n+i, 1e18) // 如果不割富人->T，那么公狗必须在 T 中
			}
		}
		ans += w // -k*fg + w+k*fg = w   先假定好友都不满足
		Fscan(in, &k)
		if tar == 0 {
			addEdge(st, n+i, w+k*fg) // 不割就表示选
		} else {
			addEdge(n+i, end, w+k*fg)
		}
	}

	dis := make([]int, len(g))
	bfs := func() bool {
		clear(dis)
		dis[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && dis[w] == 0 {
					dis[w] = dis[v] + 1
					q = append(q, w)
				}
			}
		}
		return dis[end] > 0
	}
	iter := make([]int, len(g))
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == end {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && dis[w] > dis[v] {
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
		maxFlow += dfs(st, 1e18)
	}
	Fprint(out, ans-maxFlow)
}

//func main() { cf311E(bufio.NewReader(os.Stdin), os.Stdout) }
