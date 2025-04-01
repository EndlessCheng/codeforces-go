package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf87D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	type edge struct{ v, w, wt, i int }
	g := make([]edge, n-1)
	for i := range g {
		Fscan(in, &g[i].v, &g[i].w, &g[i].wt)
		g[i].i = i + 1
	}
	slices.SortFunc(g, func(a, b edge) int { return a.wt - b.wt })

	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
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

	mx := 0
	ans := []any{}
	for _, e := range g {
		v, w := find(e.v), find(e.w)
		s := sz[v] * sz[w]
		if s > mx {
			mx = s
			ans = []any{e.i}
		} else if s == mx {
			ans = append(ans, e.i)
		}
		fa[v] = w
		sz[w] += sz[v]
	}
	Fprintln(out, mx*2, len(ans))
	Fprintln(out, ans...)
}

//func main() { cf87D(bufio.NewReader(os.Stdin), os.Stdout) }
