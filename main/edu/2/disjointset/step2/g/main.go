package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v, w, wt, ans int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &v, &w, &wt)
		es[i] = edge{v - 1, w - 1, wt}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	for _, e := range es {
		if fv, fw := f(e.v), f(e.w); fv != fw {
			ans = e.wt
			fa[fv] = fw
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
