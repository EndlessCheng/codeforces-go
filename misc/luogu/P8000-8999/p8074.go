package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p8074(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	type pair struct{ v, i int }
	a := make([]pair, n)
	b := make([]pair, n)
	c := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v, &b[i].v, &c[i].v)
		a[i].i = i
		b[i].i = i
		c[i].i = i
	}
	f := func(a, b pair) int { return a.v - b.v }
	slices.SortFunc(a, f)
	slices.SortFunc(b, f)
	slices.SortFunc(c, f)

	type edge struct{ v, w, wt int }
	es := make([]edge, 0, n*3-3)
	for i := 1; i < n; i++ {
		es = append(es, edge{a[i-1].i, a[i].i, a[i].v - a[i-1].v})
		es = append(es, edge{b[i-1].i, b[i].i, b[i].v - b[i-1].v})
		es = append(es, edge{c[i-1].i, c[i].i, c[i].v - c[i-1].v})
	}
	slices.SortFunc(es, func(a, b edge) int { return a.wt - b.wt })

	fa := make([]int, n)
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
		v := find(e.v)
		w := find(e.w)
		if v != w {
			fa[v] = w
			ans += e.wt
		}
	}
	Fprint(out, ans)
}

//func main() { p8074(bufio.NewReader(os.Stdin), os.Stdout) }
