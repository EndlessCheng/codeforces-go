package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, op int
	Fscan(in, &n, &m, &q)

	const inf int = 1e18
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i {
				g[i][j] = inf
			}
		}
	}

	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		es[i] = edge{v, w, wt}
		g[v][w] = wt
		g[w][v] = wt
	}

	type query struct{ v, w int }
	qs := make([]query, q)
	for i := range qs {
		Fscan(in, &op, &qs[i].v)
		if op == 2 {
			Fscan(in, &qs[i].w)
		} else {
			e := es[qs[i].v-1]
			g[e.v][e.w] = inf
			g[e.w][e.v] = inf
		}
	}

	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
	addEdge := func(v, w, wt int) {
		for i := range g {
			for j := range g {
				g[i][j] = min(g[i][j], g[i][v]+wt+g[w][j], g[i][w]+wt+g[v][j])
			}
		}
	}

	ans := []int{}
	for i := q - 1; i >= 0; i-- {
		q := qs[i]
		if q.w == 0 {
			e := es[q.v-1]
			addEdge(e.v, e.w, e.wt)
		} else {
			d := g[q.v-1][q.w-1]
			if d == inf {
				d = -1
			}
			ans = append(ans, d)
		}
	}
	for i := len(ans) - 1; i >= 0; i-- {
		Fprintln(out, ans[i])
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
