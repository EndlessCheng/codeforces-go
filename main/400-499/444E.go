package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf444E(in io.Reader, out io.Writer) {
	var n, tot int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 0)
		return
	}

	type edge struct{ v, w, wt int }
	es := make([]edge, n-1)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w, &es[i].wt)
	}
	slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

	fa := make([]int, n)
	sz := make([]int, n)
	lim := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
		Fscan(in, &lim[i])
		tot += lim[i]
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
		v, w := find(e.v-1), find(e.w-1)
		fa[w] = v
		sz[v] += sz[w]
		lim[v] += lim[w]
		if sz[v] > tot-lim[v] {
			Fprint(out, e.wt)
			return
		}
	}
	Fprint(out, es[n-2].wt)
}

//func main() { cf444E(bufio.NewReader(os.Stdin), os.Stdout) }
