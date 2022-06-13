package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1213G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q int
	Fscan(in, &n, &q)
	es := make([]struct{ v, w, wt int }, n-1)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })
	qs := make([]struct{ v, i int }, q)
	for i := range qs {
		Fscan(in, &qs[i].v)
		qs[i].i = i
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].v < qs[j].v })

	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	ans := make([]int64, q)
	c, i := int64(0), 0
	for _, q := range qs {
		for ; i < n-1 && es[i].wt <= q.v; i++ {
			v, w := find(es[i].v), find(es[i].w)
			c += int64(sz[v]) * int64(sz[w])
			sz[w] += sz[v]
			fa[v] = w
		}
		ans[q.i] = c
	}
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1213G(os.Stdin, os.Stdout) }
