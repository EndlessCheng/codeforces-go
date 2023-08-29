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
	var n, m, ans int
	Fscan(in, &n, &m)
	es := make([]struct{ v, w, wt int }, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	sort.Slice(es, func(i, j int) bool { return es[i].wt < es[j].wt })

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

	for _, e := range es {
		x, y := find(e.v), find(e.w)
		if x != y {
			fa[x] = y
		} else if e.wt > 0 {
			ans += e.wt
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
