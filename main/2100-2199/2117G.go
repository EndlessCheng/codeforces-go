package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2117G(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		type edge struct{ v, w, wt int }
		es := make([]edge, m)
		for i := range es {
			Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
		}
		slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

		fa := make([]int, n+1)
		mn := make([]int, n+1)
		for i := range fa {
			fa[i] = i
			mn[i] = 1e9
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
			if v == w {
				continue
			}
			mn[w] = min(mn[w], mn[v], e.wt)
			fa[v] = w
			if find(1) == find(n) {
				Fprintln(out, mn[fa[1]]+e.wt)
				break
			}
		}
	}
}

//func main() { cf2117G(bufio.NewReader(os.Stdin), os.Stdout) }
