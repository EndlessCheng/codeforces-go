package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf875F(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	type edge struct{ v, w, wt int }
	es := make([]edge, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	slices.SortFunc(es, func(a, b edge) int { return b.wt - a.wt })

	ring := make([]bool, n+1)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	for _, e := range es {
		v, w := find(e.v), find(e.w)
		if ring[v] && ring[w] {
			continue
		}
		ans += e.wt
		ring[w] = ring[w] || ring[v] || v == w
		fa[v] = w
	}
	Fprint(out, ans)
}

//func main() { cf875F(bufio.NewReader(os.Stdin), os.Stdout) }
