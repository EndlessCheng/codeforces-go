package main

import (
	. "fmt"
	"io"
	"math"
	"slices"
)

// https://github.com/EndlessCheng
func cf628F(in io.Reader, out io.Writer) {
	var n, maxV, q int
	Fscan(in, &n, &maxV, &q)
	q += 2
	type pair struct{ r, num int }
	a := make([]pair, q)
	for i := 1; i < q-1; i++ {
		Fscan(in, &a[i].r, &a[i].num)
	}
	a[q-1] = pair{maxV, n}
	slices.SortFunc(a, func(a, b pair) int { return a.r - b.r })

	const k = 5
	st := q + k
	end := st + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}

	for i := 1; i < q; i++ {
		num := a[i].num - a[i-1].num
		l, r := a[i-1].r, a[i].r
		if num < 0 || num > r-l {
			Fprint(out, "unfair")
			return
		}
		addEdge(st, i, num)
		for j := range k {
			addEdge(i, j+q, (r+k-j)/k-(l+k-j)/k)
		}
	}
	for j := range k {
		addEdge(j+q, end, n/k)
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
		maxFlow += dfs(st, math.MaxInt)
	}

	if maxFlow == n {
		Fprint(out, "fair")
	} else {
		Fprint(out, "unfair")
	}
}

//func main() { cf628F(bufio.NewReader(os.Stdin), os.Stdout) }
