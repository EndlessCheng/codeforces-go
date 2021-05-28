package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1408E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type edge struct{ v, w, wt int }
	var m, n, k, w int
	Fscan(in, &m, &n)
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	es := []edge{}
	for i, v := range a {
		for Fscan(in, &k); k > 0; k-- {
			Fscan(in, &w)
			es = append(es, edge{i, m + w - 1, v + b[w-1]})
		}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt > es[j].wt })

	fa := make([]int, m+n)
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
	ans := int64(0)
	for _, e := range es {
		if fv, fw := find(e.v), find(e.w); fv != fw {
			fa[fv] = fw
		} else {
			ans += int64(e.wt)
		}
	}
	Fprint(out, ans)
}

//func main() { CF1408E(os.Stdin, os.Stdout) }
