package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, wt int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	e1 := make([]edge, n)
	for i := range e1 {
		Fscan(in, &wt)
		e1[i] = edge{i + 1, n + 1, wt}
	}
	e2 := make([]edge, n)
	for i := range e2 {
		Fscan(in, &wt)
		e2[i] = edge{i + 1, n + 2, wt}
	}
	e3 := make([]edge, m, m+n*2)
	for i := range e3 {
		Fscan(in, &e3[i].v, &e3[i].w, &e3[i].wt)
	}

	ans := int(1e18)
	f := func(n int, edges []edge) {
		edges = append([]edge{}, edges...)
		sort.Slice(edges, func(i, j int) bool { return edges[i].wt < edges[j].wt })

		fa := make([]int, n+1)
		for i := range fa {
			fa[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			if fa[x] != x {
				fa[x] = find(fa[x])
			}
			return fa[x]
		}
		sum, cntE := 0, 0
		for _, e := range edges {
			if fv, fw := find(e.v), find(e.w); fv != fw {
				fa[fv] = fw
				sum += e.wt
				cntE++
			}
		}

		if len(edges) == m && cntE < n-1 {
			return
		}
		if sum < ans {
			ans = sum
		}
	}
	f(n, e3)
	f(n+2, append(e3, e1...))
	f(n+2, append(e3, e2...))
	f(n+2, append(append(e3, e1...), e2...))
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
